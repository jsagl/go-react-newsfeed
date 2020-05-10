package scrapped

import (
	"encoding/json"
	"errors"
	"github.com/jsagl/newsfeed-go-server/app/env"
	"github.com/jsagl/newsfeed-go-server/app/models"
	"github.com/jsagl/newsfeed-go-server/app/storage"
	"github.com/jsagl/newsfeed-go-server/lib/scrappers"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

type ExternalArticleStore struct {
	env *env.Env
}

func NewExternalArticleStore(env *env.Env) storage.ArticleStore {
	return &ExternalArticleStore{env: env}
}

func (store *ExternalArticleStore) Index() ([]*models.Article, error) {
	shouldScrap, err := store.env.Cache.Get("should_scrap")

	if err != nil {
		return nil, err
	}

	if string(shouldScrap) != "false" {
		freq, err := strconv.ParseInt(os.Getenv("SCRAPPING_FREQUENCY_IN_MIN"), 10, 64)
		if err != nil {
			return nil, err
		}

		store.env.Cache.Set("should_scrap", "false", time.Duration(freq) * time.Minute)

		go store.scrapExternalSources()
	}

	cache, err := store.env.Cache.Get("scrapped_articles")

	if err != nil {
		return nil, err
	}

	if cache != nil {
		articles, err := store.parseCache(cache)
		if err != nil {
			return nil, err
		}

		return articles, nil
	}

	mockArticle := &models.Article{
		Title: "Sources will be available shortly. Please refresh the page",
		TargetUrl: "https://www.rubyflow.com",
		Date: time.Now(),
		Category: "golang",
		SourceName: "Go React Newsfeed",
		SourceUrl: "https://www.rubyflow.com",
	}

	return []*models.Article{mockArticle}, nil
}

func (store *ExternalArticleStore) scrapExternalSources() {
	store.env.Logger.Infow("started scrapping")

	var wg sync.WaitGroup

	sources := []string{"ruby_flow", "ruby_weekly", "thoughtbot", "golang_weekly"}

	jobs := make(chan string, len(sources))
	articlesChan := make(chan []*models.Article, len(sources))


	for worker := 1; worker <= 5; worker++ {
		wg.Add(1)
		go store.sourceScrapperWorker(jobs, articlesChan, &wg, worker)
	}

	for _, source := range sources {
		jobs <- source
	}

	close(jobs)

	wg.Wait()

	close(articlesChan)

	articles := make([]*models.Article, 0)
	for articlesFromSource := range articlesChan {
		articles = append(articles, articlesFromSource...)
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Date.After(articles[j].Date)
	})

	store.storeInCache(articles)
}

func (store *ExternalArticleStore) sourceScrapperWorker(jobs <-chan string, articlesChan chan<- []*models.Article, wg *sync.WaitGroup, worker int) {
	defer wg.Done()

	for source := range jobs {
		scrappedArticles, err := store.scrapSource(source)
		if err != nil {
			store.env.Logger.Infow("failed_to_scrap_ressource", "source", source, "error", err)
		} else {
			articles := make([]*models.Article, 0)
			for _, scrappedArticle := range scrappedArticles {
				article := store.convertToArticleModel(scrappedArticle)
				articles = append(articles, article)
			}
			articlesChan <- articles
		}
	}

	return
}

func (store *ExternalArticleStore) scrapSource(source string) ([]scrappers.Article, error) {
	switch source {
	case "ruby_flow":
		articles, err := scrappers.ScrapRubyflowArticles()
		if err != nil {
			return nil, err
		}
		return articles, nil

	case "ruby_weekly":
		articles, err := scrappers.ScrapRubyWeeklyArticles()
		if err != nil {
			return nil, err
		}
		return articles, nil

	case "golang_weekly":
		articles, err := scrappers.ScrapGolangWeeklyArticles()
		if err != nil {
			return nil, err
		}
		return articles, nil

	case "thoughtbot":
		articles, err := scrappers.ScrapThoughtbotArticles()
		if err != nil {
			return nil, err
		}
		return articles, nil
	}

	return nil, errors.New("unknown source to scrap from")
}

func (store *ExternalArticleStore) convertToArticleModel (sourceArticle scrappers.Article) *models.Article {
	return &models.Article{
		Title: sourceArticle.Title,
		TargetUrl: sourceArticle.TargetUrl,
		Date: sourceArticle.Date,
		Category: sourceArticle.Category,
		SourceUrl: sourceArticle.SourceUrl,
		SourceName: sourceArticle.SourceName,
	}
}

func (store *ExternalArticleStore) parseCache(cache []byte) ([]*models.Article, error) {
	var articles []*models.Article

	err := json.Unmarshal(cache, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (store *ExternalArticleStore) storeInCache(articles []*models.Article) {
	encodedArticles, err := json.Marshal(articles)
	if err != nil {
		store.env.Logger.Infow("failed to store scrapped articles in cache", "error", err)
	}

	store.env.Cache.Set("scrapped_articles", encodedArticles, 0)
}

