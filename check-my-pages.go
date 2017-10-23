package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/asciimoo/colly"
)

func main() {

	isHelp := flag.Bool("help", false, "Help")
	urlsFileName := flag.String("urls", "urls.csv", "Name of the csv file with the urs in the first column")
	isHTTP := flag.Bool("http", false, "Http response codes")
	isAnalytics := flag.Bool("analytics", false, "Correct analytics tag in the html")
	isCanonical := flag.Bool("canonical", false, "Canonical URLS in the ")
	isLinkpattern := flag.Bool("linkpattern", false, "Link Pattern")
	isCSSJsPattern := flag.Bool("cssjspattern", false, "CSS and JS Pattern")
	isMediaPattern := flag.Bool("mediapattern", false, "Image, object and iframe Pattern")
	pattern := flag.String("pattern", `https?://(\w|-)+.greenpeace.org/espana/.+`, "Regular expression to detect in the links")
	waitMiliseconds := flag.Int("miliseconds", 100, "Miliseconds between requests")
	isClear := flag.Bool("clear", false, "Remove files created by this script")
	flag.Parse()

	allUrlsCsv := readCsvFile(*urlsFileName)

	allUrls := csvFirstColumnToSlice(allUrlsCsv)

	linkRegex, _ := regexp.Compile(*pattern)

	c := colly.NewCollector()
	// c.AllowedDomains = []string{"localhost", "greenpeace.es", "archivo.greenpeace.es"}

	if *isHelp == true {
		help()
		os.Exit(0)
	}

	if *isHTTP == true {

		isHTTPfile, isHTTPErr := os.OpenFile("httpResponses.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if isHTTPErr != nil {
			panic(isHTTPErr)
		}
		defer isHTTPfile.Close()

		var lineHTTP string
		for _, v := range allUrls {
			lineHTTP = getHTTPinfoAsCsvline(v)
			if _, err := isHTTPfile.WriteString(lineHTTP); err != nil {
				panic(err)
			}
			time.Sleep(time.Millisecond * time.Duration(*waitMiliseconds))
		}

		os.Exit(0)
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

	if *isCSSJsPattern == true {

		cssJsPattern, cssJsPatternErr := os.OpenFile("cssjspattern.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if cssJsPatternErr != nil {
			panic(cssJsPatternErr)
		}
		defer cssJsPattern.Close()

		c.OnHTML("link[rel=stylesheet]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if linkRegex.MatchString(link) {
				lineCSSJsPattern := fmt.Sprintf("%s,css,%s\n", e.Request.URL.String(), link)
				if _, err := cssJsPattern.WriteString(lineCSSJsPattern); err != nil {
					panic(err)
				}
			}

		})

		c.OnHTML("script", func(e *colly.HTMLElement) {
			src := e.Attr("src")
			if linkRegex.MatchString(src) {
				lineCSSJsPattern := fmt.Sprintf("%s,js,%s\n", e.Request.URL.String(), src)
				if _, err := cssJsPattern.WriteString(lineCSSJsPattern); err != nil {
					panic(err)
				}
			}

		})
	}

	if *isMediaPattern == true {

		mediaPattern, mediaPatternErr := os.OpenFile("mediapattern.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if mediaPatternErr != nil {
			panic(mediaPatternErr)
		}
		defer mediaPattern.Close()

		c.OnHTML("img, picture source", func(e *colly.HTMLElement) {
			src := e.Attr("src")
			srcset := e.Attr("srcset")
			if src == "" && srcset != "" {
				src = srcset
			}
			if linkRegex.MatchString(src) {
				lineMediaPattern := fmt.Sprintf("%s,img,%s\n", e.Request.URL.String(), src)
				if _, err := mediaPattern.WriteString(lineMediaPattern); err != nil {
					panic(err)
				}
			}

		})

		c.OnHTML("video, video source", func(e *colly.HTMLElement) {
			src := e.Attr("src")
			if linkRegex.MatchString(src) {
				lineMediaPattern := fmt.Sprintf("%s,video,%s\n", e.Request.URL.String(), src)
				if _, err := mediaPattern.WriteString(lineMediaPattern); err != nil {
					panic(err)
				}
			}

		})

		c.OnHTML("audio, audio source", func(e *colly.HTMLElement) {
			src := e.Attr("src")
			if linkRegex.MatchString(src) {
				lineMediaPattern := fmt.Sprintf("%s,audio,%s\n", e.Request.URL.String(), src)
				if _, err := mediaPattern.WriteString(lineMediaPattern); err != nil {
					panic(err)
				}
			}

		})

		c.OnHTML("iframe", func(e *colly.HTMLElement) {
			src := e.Attr("src")
			if linkRegex.MatchString(src) {
				lineMediaPattern := fmt.Sprintf("%s,iframe,%s\n", e.Request.URL.String(), src)
				if _, err := mediaPattern.WriteString(lineMediaPattern); err != nil {
					panic(err)
				}
			}

		})

		c.OnHTML("object", func(e *colly.HTMLElement) {
			src := e.Attr("data")
			if linkRegex.MatchString(src) {
				lineMediaPattern := fmt.Sprintf("%s,object,%s\n", e.Request.URL.String(), src)
				if _, err := mediaPattern.WriteString(lineMediaPattern); err != nil {
					panic(err)
				}
			}

		})

	}

	if *isClear == true {

		os.Remove("httpResponses.csv")
		os.Remove("analytics.csv")
		os.Remove("canonicals.csv")
		os.Remove("linkpattern.csv")
		os.Remove("cssjspattern.csv")
		os.Remove("mediapattern.csv")
		os.Exit(0)
	}

	// Open URLs file
	for _, v := range allUrls {
		c.Visit(v)
		time.Sleep(time.Millisecond * time.Duration(*waitMiliseconds))
	}

}
