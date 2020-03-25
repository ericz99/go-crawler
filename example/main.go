package example

import (
	crawler "github.com/ericz99/go-crawler"
)

func main() {
	// # create a crawler instance
	spider := crawler.Crawler{}
	// # crawl the page
	result, domain := spider.Crawl("https://kith.com/")
	// # download result
	spider.Download(result, domain)
}
