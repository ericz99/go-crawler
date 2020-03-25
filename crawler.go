package crawler

import (
	"bufio"
	"fmt"
	"go-crawler/models"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/*
deep crawl specific websites, and get all url links + relative path

find a tag src, script tag src, link href

then download the links as site.txt

only scrape the page itself, and not other links
*/

// Crawler struct (MODEL)
type Crawler struct {
	Proxy models.Proxy `json:"proxy"`
}

// ScrapeResult struct (MODEL)
type ScrapeResult struct {
	Link string `json:"link"`
}

// Get - REQUEST METHOD
func Get(url string) (*http.Response, error) {
	// # custom http client
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	// # make request
	req, err := http.NewRequest("GET", url, nil)
	// # add headers to avoid issues with sites sending error codes for default golang user agent
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	req.Header.Add("Accept", "*/*")

	// # if request error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// # get response
	resp, err := client.Do(req)

	// # if response error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// # return response
	return resp, nil
}

// GetDomain - Method | Returns domain of the link
func GetDomain(startURL string) string {
	parsed, _ := url.Parse(startURL)
	return fmt.Sprintf(parsed.Host)
}

// GetAllLink - Method
func GetAllLink(tag string, doc *goquery.Document, c chan []ScrapeResult) {
	// # create an array
	results := []ScrapeResult{}
	// # scrape tag
	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		if _, exist := s.Attr("src"); exist {
			res, _ := s.Attr("src")
			if res != "" {
				results = append(results, ScrapeResult{Link: res})
			}
		} else {
			res, _ := s.Attr("href")
			if res != "" {
				results = append(results, ScrapeResult{Link: res})
			}
		}
	})

	// # forward result to channel
	c <- results
}

// ExtractLink - Method | return all links in the current page. Including path, link, etc
func ExtractLink(doc *goquery.Document) []ScrapeResult {
	links := make(chan []ScrapeResult)
	// # list of stuff to scrape
	list := []string{"a", "script", "link"}
	// # create an array
	results := []ScrapeResult{}
	if doc != nil {
		for _, item := range list {
			// # do logic, get all link method concurrently
			go GetAllLink(item, doc, links)
			// # get the results from channel
			allLink := <-links
			// # append result to main array
			results = append(results, allLink...)
		}
	}

	return results
}

// CrawlPage - Method | scrape all urls, paths, endpoints on the page
func CrawlPage(startURL string) []ScrapeResult {
	// # define scrape result
	results := []ScrapeResult{}
	// # make an request
	resp, err := Get(startURL)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
		return results
	}

	// # load html doc response
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// # extract all links
	links := ExtractLink(doc)
	// # append the list
	results = append(results, links...)
	// # return all links
	return results
}

// Download - Method | download the data to file system directory
func (c Crawler) Download(data []ScrapeResult, domain string) {
	fmt.Println("DOWNLOADING TO FILE!")
	str := ""
	// # download it to file
	f, err := os.Create(fmt.Sprintf("%s.txt", domain))
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	for _, d := range data {
		str = str + fmt.Sprintf("%s\n", d.Link)
	}

	n4, err := w.WriteString(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}

// Crawl - METHOD
func (c Crawler) Crawl(startURL string) ([]ScrapeResult, string) {
	// # define scrape result
	results := []ScrapeResult{}
	// # get domain url
	baseDomain := GetDomain(startURL)
	// # crawl page
	links := CrawlPage(startURL)
	// # append links to result
	results = append(results, links...)

	return results, baseDomain
}
