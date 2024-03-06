package main

import (
  "database/sql"
  "fmt"
  "log"

  "github.com/gocolly/colly"
  _ "github.com/go-sql-driver/mysql"
)

type PageLink struct {
  title string
  link string
}


func main() {
  scrapeFlag := false

  connectDataBase()

  // NOTE keep this from runnint for now
  if scrapeFlag == true {
    // send scarper to tarURL.
    tarURL := "https://9to5mac.com"
    articles := scrapeUrl(tarURL, "a.article__title-link")

    // print scraped results.
    fmt.Println("Collected links:")
    for _, article := range articles {
      fmt.Printf("\n%s \n%s\n", article.title, article.link)
    }
  }
}


// Opens and closes the onection to the database.
func connectDataBase() *sql.DB{
  // Database connection details - Generaly this should not be hard codded.
  // NOTE pass this in insted of hard codeing
  dsn := "debian-sys-maint:fAkBoMzWqEDlkGNf@tcp(localhost:3306)/go_scrape?charset=utf8mb4"

  // Connect to the database
  db, err := sql.Open("mysql", dsn)
  defer db.Close() // Close the connection when the program exits
  
  // Check if databased open retunred errer
  if err != nil {
    log.Fatal(err)
  }
  
  // Check that the database connected.
  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
  return db
}


// inserts data to dataBase
func insertToDatabase(db *sql.DB, article PageLink) {
  quary := ""
}


// TODO
func userInput() {

}


// TODO
func linkSearch(articles []PageLink, searchTarget string) {
  
}


// Scrapes a url and returns the slice of links with there titles
//  Selectore is the html element you are targetting.
func scrapeUrl(targetUrl string, selector string) []PageLink {
  // Instantiate default collector
  c := colly.NewCollector(
    //colly.AllowedDomains("9to5mac.com", "9to5mac.com"),
  )
 
  var articles []PageLink

  // On every a element which has selector attribute scrape the link and title
  c.OnHTML(selector, func(e *colly.HTMLElement) {
    link := e.Attr("href")
    title := e.Text

    page := PageLink{link:link, title: title}
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

