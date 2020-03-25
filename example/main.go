package main

import (
	crawler "github.com/ericz99/crawler"
)

func main() {
	// # create a crawler instance with option
	spider := crawler.Crawler{Option: crawler.Option{Concurrency: 10}}
	// # crawl the page
	result, domain := spider.Crawl("https://kith.com/")
	// # download result
	spider.Download(result, domain)
}
