package main

import (
	"fmt"
)

func help() {
	fmt.Println(`

-------------------
CHECK-MY-PAGES HELP
-------------------

check-my-pages is a scrapping script. It checks each url in a list and creates report files about what was tested. Each file reports about a specific issue and includes the scanned url together with the result.


EXAMPLES:

./check-my-pages -urls=urls.csv -http -miliseconds=100

./check-my-pages -urls=urls.csv -analytics -canonical -title -linkpattern -cssjspattern -mediapattern


CHECKS:

-http : Gets the http response code, mime-type, file size and final url. It must be used separately from the other checks.

-analytics : Gets the first Google Analytics account.

-canonical : Gets the canonical URL for the url.

-title : Gets the title for the url, if it's an html page

-linkpattern : Gets links that match the regular expression pattern.

-cssjspattern : Gets CSS and JS URLs that match the regular expression pattern.

-mediapattern : Gets urls from images, videos, audios, iframes and objects that match the regular expression pattern


OPTIONS:

-urls=urls.csv : Sets the file with the urls to scan. Normally a text file with one URL per line or a csv without headers with the urls on the first column.

-pattern='https?://(\w|-)+.greenpeace.org/espana/.+' : Changes the url search pattern to the regular expression. To be used with *pattern checks.

-miliseconds=100 : Sets a delay of 100 miliseconds between requests.

OTHER:

-clear : Deletes all the files with the reports

-stash : Renames all the files with the reports to prevent overwiting them


FILES WITH THE REPORTS:

- httpResponses.csv : Stores the http response codes for the URL. 200 means everything is OK.

- analytics.csv : Reports google analytics tracking ID.

- canonicals.csv : Reports the canonical url for every url

- titles.csv : Reports the title for every url

- linkpattern.csv : Reports on links that include a regular expression pattern. Useful to track links to specific dead sites. The default pattern can be set by the -pattern option.

- cssjspattern.csv : Reports css and js urls that include a regular expression pattern. To detect dead css and js urls in large sites. The pattern can also be defined with the option -pattern (described bellow)

- mediapattern.csv : Reports media links. Images, videos, audios, iframes and objects. Also use -pattern to define the urls pattern.

CRAWL

If you don't have a file with the urls you can try to obtain it by crawling the site:

./check-my-pages -crawl -urls=crawledurls.csv  -start='https://www.fotografar.net/' -pattern='https://www.fotografar.net/.*'

It will save, in crawledurls.csv, all the urls it can find from the start url and that match the pattern.

	`)
}
