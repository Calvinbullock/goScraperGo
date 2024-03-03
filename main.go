package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type PageLinks struct {
  title string
  link string
}

func main() {
  tarURL := "https://9to5mac.com"
  articles := scrapeUrl(tarURL, "a.article__title-link")
 
  fmt.Println("Collected links:")
  for _, article := range articles {
    fmt.Printf("\n%s \n%s\n", article.title, article.link)
  }
}

func linkSearch(articles []PageLinks) {

}

// Scrapes a url and returns the slice of links with there titles
func scrapeUrl(targetUrl string, selector string) []PageLinks {
  // Instantiate default collector
  c := colly.NewCollector(
    //colly.AllowedDomains("9to5mac.com", "9to5mac.com"),
  )
 
  var articles []PageLinks

  // On every a element which has selector attribute scrape the link and title
  c.OnHTML(selector, func(e *colly.HTMLElement) {
    link := e.Attr("href")
    title := e.Text

    page := PageLinks{link:link, title: title}
    articles = append(articles, page)
  })

  // Start scraping on tarURL
  err := c.Visit(targetUrl)
  
  //
  if err != nil {
    fmt.Println("Error:", err)
  }
  
  // statues print
  fmt.Printf("\nc: %v\n", c)
  return articles 
}

