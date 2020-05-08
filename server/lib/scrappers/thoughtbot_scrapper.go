package scrappers

import (
	"github.com/PuerkitoBio/goquery"
	"time"
)

func ScrapThoughtbotArticles() ([]Article, error) {
	doc, err := RequestDocument("https://www.thoughtbot.com/blog")
	if err != nil {
		return nil, err
	}

	articles := make([]Article, 0)
	dateLayout := "2006-01-02T15:04:05-07:00"

	doc.Find(".mini-post").Each(func(i int, s *goquery.Selection) {
		var title, targetUrl, category, sourceUrl, sourceName  string
		var date time.Time

		title = s.Find(".mini-post-link").First().Text()
		category = "thoughtbot"
		sourceUrl = "https://www.thoughtbot.com/blog"
		sourceName = "Thoughtbot"

		datestr, exists := s.Find(".meta-date").Attr("datetime")
		if exists {
			date, _ = time.Parse(dateLayout, datestr)

		}

		urlComplement, exists := s.Find(".mini-post-link").First().Attr("href")
		if exists {
			targetUrl = "https://www.thoughtbot.com" + urlComplement
		}

		article := Article{
			Title: title,
			Category: category,
			SourceUrl: sourceUrl,
			SourceName: sourceName,
			Date: date,
			TargetUrl: targetUrl,
		}

		articles = append(articles, article)
	})

	return articles, nil
}