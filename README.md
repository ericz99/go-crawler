## go-crawler

Will crawl through the website, and scrape all endpoints, paths, hashtags, etc.

## Installation

Installation is done using `go get`.

```
go get -u github.com/ericz99/go-crawler
```

## Example

```golang
package example

import (
	crawler "github.com/ericz99/go-crawler"
)

func main() {
	// # create a crawler instance with option
	spider := crawler.Crawler{}
	// # crawl the page
	result, domain := spider.Crawl("https://kith.com/")
	// # download result
	spider.Download(result, domain)
}
```

## Todo

- [ ] Find all links base on regex
- [ ] Add proxy support
