// Exercise: Web Crawler
// In this exercise you'll use Go's concurrency features to parallelize a web crawler.

// Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Safe cache is ok to use concurrently.
// We need to make sure that only one goroutine can access this
// cache at a time to avoid conflicts.
type SafeCache struct {
	v   map[string]bool
	mux sync.Mutex
}

func (c *SafeCache) CacheUrl(url string) {
	c.mux.Lock()
	c.v[url] = true
	c.mux.Unlock()
}

func (c *SafeCache) Value(url string) bool {
	c.mux.Lock()
	// unlock the map after the value is returned
	defer c.mux.Unlock()
	return c.v[url]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache) {
	// TODO: Don't fetch the same URL twice.
	// We need to have a cache object that lets us know whether
	// we have fetched the url before.

	// If we have already fetched this url before,
	// do not fetch it again.
	if cache.Value(url) {
		fmt.Println("\nAlready visited ", url)
		return
	}

	if depth <= 0 {
		fmt.Printf("\nDone crawling, %v \n\n", url)
		return
	}

	fmt.Printf("\nCrawling %v", url)
	// get the urls found from the given url's body
	_, childUrls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	cache.CacheUrl(url)

	// create a channel to see whether each child url
	// is done fetching urls
	done := make(chan bool)

	// - for each of the urls, we want to crawl the url
	// to find the urls from that page (start a new goroutine)
	// - decrement depth by 1 since we are looking at the child url
	for _, u := range childUrls {
		go func(childUrl string) {
			Crawl(childUrl, depth-1, fetcher, cache)
			cache.CacheUrl(childUrl)
			done <- true
		}(u)
	}

	// Wait to receive the done boolean from the child url
	// Here we are receiving the boolean value from the channel
	for range childUrls {
		<-done
		fmt.Printf("\nDone with %v", url)
	}
	return
}

func main() {
	cache := SafeCache{v: make(map[string]bool)}

	Crawl("https://golang.org/", 4, fetcher, &cache)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
