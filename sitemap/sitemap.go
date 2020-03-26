package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/ambye85/gophercises/link"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls []loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}

//func bfs(urlStr string, maxDepth int) []string {
//	seen := make(map[string]struct{})
//	var q map[string]struct{}
//	nq := map[string]struct{}{
//		urlStr: struct{}{},
//	}
//	for i := 0; i <= maxDepth; i++ {
//		q, nq = nq, make(map[string]struct{})
//		for href, _ := range q {
//			if _, ok := seen[href]; ok {
//				continue
//			}
//			seen[href] = struct{}{}
//			for _, l := range get(href) {
//				nq[l] = struct{}{}
//			}
//		}
//	}
//	ret := make([]string, 0, len(seen))
//	for u, _ := range seen {
//		ret = append(ret, u)
//	}
//	return ret
//}

type empty struct{}

func bfs(urlStr string, maxDepth int) []string {
	var results []string
	visitedPages := make(map[string]empty)
	toBeVisited := []string{urlStr}
	currentDepth := 0
	for len(toBeVisited) != 0 && currentDepth <= maxDepth {
		page := toBeVisited[0]
		if _, visited := visitedPages[page]; !visited {
			visitedPages[page] = empty{}
			results = append(results, page)
			for _, l := range get(page) {
				toBeVisited = append(toBeVisited, l)
			}
		}
		toBeVisited = toBeVisited[1:]
	}
	return results
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(body io.Reader, base string) []string {
	links, _ := link.Parse(body)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, fltr func(string) bool) []string {
	var ret []string
	for _, l := range links {
		if fltr(l) {
			ret = append(ret, l)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(l string) bool {
		return strings.HasPrefix(l, pfx)
	}
}

//func main() {
//	// if link not already visited
//	// Add link to queue
//	// Get page
//	// Parse links
//
//	rootUrl := flag.String("url", "https://google.com/", "the root url from which to build the sitemap")
//
//	if !strings.HasSuffix(*rootUrl, "/") {
//		*rootUrl = *rootUrl + "/"
//	}
//
//	visitedUrls := map[string]bool{*rootUrl: true}
//	nextUrls := []string{*rootUrl}
//
//	for page := nextUrls[0]; len(nextUrls) != 0; nextUrls = nextUrls[1:] {
//		//page := nextUrls[0]
//		//nextUrls = nextUrls[1:]
//
//		visitedUrls[page] = true
//		resp, err := http.Get(page)
//		if err != nil {
//			panic(err)
//		}
//		defer resp.Body.Close()
//		links, err := link.Parse(resp.Body)
//		if err != nil {
//			panic(err)
//		}
//
//		for _, url := range links {
//			href := url.Href
//			if strings.HasPrefix(href, "/") {
//				href = *rootUrl + href[1:]
//			}
//			if _, visited := visitedUrls[url.Href]; !visited && (strings.HasPrefix(url.Href, *rootUrl)) {
//				nextUrls = append(nextUrls, url.Href)
//				fmt.Printf("%+v", visitedUrls)
//				fmt.Println(!visited, url.Href)
//			}
//		}
//	}
//}
