package scrappers

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func ScrapJSWeeklyArticles() ([]Article, error) {
	scrappingError := "Error while scrapping JSWeekly"

	// Necessary to skip certificates validation
	tr:= &http.Transport {TLSClientConfig:&tls.Config {InsecureSkipVerify: true}}
	client := http.Client {Transport: tr}

	resp1, err := client.Get ("https://javascriptweekly.com/")
	if err != nil {
		return nil, err
	}
	defer resp1.Body.Close()

	homepage, err := goquery.NewDocumentFromReader(resp1.Body)
	if err != nil {
		return nil, err
	}

	urlComplement, ok := homepage.Find(".main").Find("p").First().Find("a").Attr("href")
	if !ok {
		fmt.Println("could not find latest newsletter issue number at JSWeekly")
		return nil, nil
	}

	url := "https://javascriptweekly.com" + urlComplement
	resp2, err := client.Get (url)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp2.Body)
	if err != nil {
		return nil, err
	}

	date, err := extractDateJSWeekly(urlComplement)
	if err != nil {
		fmt.Println(scrappingError)
		return nil, err
	}

	articles := make([]Article, 0)

	doc.Find("#content").Find(".el-item").Each(func(i int, s *goquery.Selection) {
		var isSponsored bool
		isSponsored = s.Find("span").Last().HasClass("tag-sponsor")

		if !isSponsored {
			var title, targetUrl, category, sourceUrl, sourceName  string

			title = s.Find(".mainlink").First().Find("a").Text()
			category = "javascript"
			sourceUrl = "https://javascriptweekly.com/"
			sourceName = "JavascriptWeekly"
			targetUrl, _ = s.Find(".mainlink").First().Find("a").Attr("href")

			article := Article{
				Title: title,
				Category: category,
				SourceUrl: sourceUrl,
				SourceName: sourceName,
				Date: date,
				TargetUrl: targetUrl,
			}

			articles = append(articles, article)
		}
	})

	return articles, nil
}

func extractDateJSWeekly(urlComplement string) (time.Time, error) {
	r, _ := regexp.Compile(`\d+`)
	issueNumberAsString := r.FindString(urlComplement)

	var issueNumber int64
	var err error

	if issueNumberAsString != "" {
		issueNumber, err = strconv.ParseInt(issueNumberAsString, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
	}

	referenceUnixTime := int64(1588921200)
	referenceIssueNum := int64(487)
	oneWeekInSec := int64(604800)

	unixTime := referenceUnixTime + (issueNumber - referenceIssueNum) * oneWeekInSec
	date := time.Unix(unixTime, 0)

	return date, nil
}