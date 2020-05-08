package scrappers

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"time"
)

func ScrapRubyflowArticles() ([]Article, error) {
	doc, err := RequestDocument("https://www.rubyflow.com/")
	if err != nil {
		return nil, err
	}

	articles := make([]Article, 0)

	doc.Find(".posts").Find(".post").Not(".sponsored").Each(func(i int, s *goquery.Selection) {
		var title, targetUrl, category, sourceUrl, sourceName  string
		var date time.Time

		title = s.Find("h1").First().Find("a").Text()
		category = "ruby"
		sourceUrl = "https://www.rubyflow.com/"
		sourceName = "RubyFlow"

		timestamp, exists := s.Attr("data-timestamp")
		if exists {
			i, err := strconv.ParseInt(timestamp, 10, 64)
			if err == nil {
				date = time.Unix(i, 0)
			}
		}

		link := s.Find("p").First().Find("a").First()

		if link.Parent().HasClass("more") {
			ref, _ := link.Attr("href")
			visit := "https://www.rubyflow.com" + ref
			subdoc, err := RequestDocument(visit)
			if err == nil {
				url, exists := subdoc.Find(".content").Find("a").First().Attr("href")
				if exists {
					targetUrl = url
				}
			}
		} else {
			url, exists := link.Attr("href")
				if exists {
					targetUrl = url
				}
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