
package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

type PageInfo struct {
  title string
  link string
}

func main() {
  //tarURL := "https://www.thecoupleproject.com"
  //tarURL := "https://www.amazon.com"
  tarURL := "https://9to5mac.com"

  // Instantiate default collector
  c := colly.NewCollector(
    // Visit only domains: hackerspaces.org, wiki.hackerspaces.org
    //colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
    //colly.AllowedDomains("www.thecoupleproject.com", "thecoupleproject.com"),
    colly.AllowedDomains("9to5mac.com", "9to5mac.com"),
  )
 
  var articles []PageInfo 
  element := "h[ref]"

  // On every a element which has href attribute call callback
  c.OnHTML(element, func(e *colly.HTMLElement) {
    link := e.Attr("href")
    // Print link
    fmt.Printf("Link found: %q -> %s\n", e.Text, link)

    page := PageInfo {link:link}
    articles = append(articles, page)

    // Visit link found on page
    // Only those links are visited which are in AllowedDomains
    //c.Visit(e.Request.AbsoluteURL(link))
  })
  
  /*
  // Before making a request print "Visiting ..."
  c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting", r.URL.String())
  })
  */
  c.OnError(func(r *colly.Response, e error) {
    fmt.Printf("Error on scrap: %s\n", e.Error())
  })

  // Start scraping on tarURL
  c.Visit(tarURL)

  for _, article := range articles {
    fmt.Println(article)
  }

  fmt.Printf("c: %v\n", c)

}

