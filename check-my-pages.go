package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/asciimoo/colly"
)

func main() {

	urlsFileName := flag.String("urls", "urls.csv", "Name of the csv file with the urs in the first column")
	isHTTP := flag.Bool("http", false, "Http response codes")
	isRedirects := flag.Bool("redirects", false, "Redirects response codes")
	isAnalytics := flag.Bool("analytics", false, "Correct analytics tag in the html")
	isCanonical := flag.Bool("canonical", false, "Canonical URLS in the ")
	isLinkpattern := flag.Bool("linkpattern", false, "Link Pattern")
	waitMiliseconds := flag.Int("miliseconds", 100, "Miliseconds between requests")
	isClear := flag.Bool("clear", false, "Remove files created by this script")
	flag.Parse()

	allUrlsCsv := readCsvFile(*urlsFileName)

	allUrls := csvFirstColumnToSlice(allUrlsCsv)

	linkRegex, _ := regexp.Compile(`https?://(\w|-)+.greenpeace.org/espana/.+`)

	c := colly.NewCollector()
	// c.AllowedDomains = []string{"localhost", "greenpeace.es", "archivo.greenpeace.es"}

	if *isHTTP == true {

		httpResponses, httpErr := os.OpenFile("httpResponses.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if httpErr != nil {
			panic(httpErr)
		}
		defer httpResponses.Close()

		c.OnResponse(func(r *colly.Response) {
			lineResponse := fmt.Sprintf("%s,%v\n", r.Request.URL.String(), r.StatusCode)
			if _, err := httpResponses.WriteString(lineResponse); err != nil {
				panic(err)
			}

		})
	}

	if *isAnalytics == true {

		analytics, analyticsErr := os.OpenFile("analytics.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if analyticsErr != nil {
			panic(analyticsErr)
		}
		defer analytics.Close()

		c.OnResponse(func(r *colly.Response) {
			body := string(r.Body)
			foundUA := searchInString(body, `UA-\d{5,8}-\d{1,2}`)
			lineResponse := fmt.Sprintf("%s,%s\n", r.Request.URL.String(), foundUA)
			if _, err := analytics.WriteString(lineResponse); err != nil {
				panic(err)
			}
		})
	}

	if *isCanonical == true {

		canonical, canonicalErr := os.OpenFile("canonicals.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if canonicalErr != nil {
			panic(canonicalErr)
		}
		defer canonical.Close()

		c.OnHTML("link[rel=canonical]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			lineCanonical := fmt.Sprintf("%s,%s\n", e.Request.URL.String(), link)
			if _, err := canonical.WriteString(lineCanonical); err != nil {
				panic(err)
			}
		})
	}

	if *isLinkpattern == true {

		linkpattern, linkpatternErr := os.OpenFile("linkpattern.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if linkpatternErr != nil {
			panic(linkpatternErr)
		}
		defer linkpattern.Close()

		c.OnHTML("a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if linkRegex.MatchString(link) {
				lineLinkpattern := fmt.Sprintf("%s,%s\n", e.Request.URL.String(), link)
				if _, err := linkpattern.WriteString(lineLinkpattern); err != nil {
					panic(err)
				}
			}

		})
	}

	if *isRedirects == true {

		redirects, redirectsErr := os.OpenFile("redirects.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if redirectsErr != nil {
			panic(redirects)
		}
		defer redirects.Close()

		c.OnRequest(func(r *colly.Request) {
			response, error := http.Get(r.URL.String())
			if error != nil {
				fmt.Printf("=> %v\n", error.Error())
			} else {
				finalURL := response.Request.URL.String()
				lineCanonical := fmt.Sprintf("%s,%s\n", r.URL.String(), finalURL)
				if _, err := redirects.WriteString(lineCanonical); err != nil {
					panic(err)
				}
			}

		})
	}

	if *isClear == true {

		os.Remove("httpResponses.csv")
		os.Remove("analytics.csv")
		os.Remove("canonicals.csv")
		os.Remove("redirects.csv")
		os.Exit(0)
	}

	// Open URLs file
	for _, v := range allUrls {
		c.Visit(v)
		time.Sleep(time.Millisecond * time.Duration(*waitMiliseconds))
	}

}
