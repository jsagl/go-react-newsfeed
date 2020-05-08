package scrappers

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

type Article struct {
	Title   	string		`json:"title,omitempty"`
	TargetUrl 	string		`json:"target_url,omitempty"`
	Date		time.Time	`json:"date,omitempty"`
	Category	string		`json:"category,omitempty"`
	SourceUrl	string		`json:"source_url,omitempty"`
	SourceName	string		`json:"source_url,omitempty"`
}

func RequestDocument(url string) (*goquery.Document, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return doc, nil
}

func deleteEmptyStringsFromSlice (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func cleanString(s string) string {
	str := strings.Replace(s, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "  ", "", -1)

	return str
}