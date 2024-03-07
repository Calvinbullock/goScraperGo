package main

import (
  "database/sql"
  "fmt"
  "log"

  "github.com/gocolly/colly"
  _ "github.com/go-sql-driver/mysql"
)

type PageLink struct {
  id int
  title string
  link string
}

func main() {
  // User input for scraping or not.
  var scrapeFlag rune
  fmt.Println("Do you want to scrape? (y/n): ")
  fmt.Scanf("%s\n", scrapeFlag)

  // NOTE Database connection details - Generaly this should not be hard codded.
  dbCredentals := "debian-sys-maint:fAkBoMzWqEDlkGNf@tcp(localhost:3306)/go_scrape?charset=utf8mb4"
  db := connectDataBase(dbCredentals)

  if scrapeFlag == 0 {
    // send scarper to tarURL.
    tarURL := "https://9to5mac.com"
    articles := scrapeUrl(tarURL, "a.article__title-link")

    // send scraped results to DataBase.
    fmt.Println("Collected links:")
    for _, article := range articles {
      fmt.Printf("\n%d \n%s \n%s\n", article.id, article.title, article.link)
      insertToDatabase(db, article)
    }
  }
  // Close the DataBase connection when the program exits
  defer db.Close() 
}

// Opens and closes the conection to the database.
func connectDataBase(dbCredentals string) *sql.DB{
  // Connect to the database
  db, err := sql.Open("mysql", dbCredentals)

  // Check if databased open returned errer
  if err != nil {
    log.Fatal("DB open? ", err)
  }

  // Check that the database connected.
  if err = db.Ping(); err != nil {
    log.Fatal("DB ping? ", err)
  }
  return db
}

// inserts data to dataBase
func insertToDatabase(db *sql.DB, article PageLink) {
  quary := "INSERT INTO article_table(id, title, link) VALUES (?, ?, ?)"

  err := db.QueryRow(quary, article.id, article.title, article.link)
  if err != nil {
    fmt.Println("Insert fault: ", err)
  }
}

// TODO
func userInput() {

}

// TODO
func linkSearch(articles []PageLink, searchTarget string) {

}

// Scrapes a url and returns a slice of PageLinks.
//  Selector is the html element you are targetting.
func scrapeUrl(targetUrl string, selector string) []PageLink {
  // id for when the an article is added to the database.
  id := 0
  var articles []PageLink

  // Instantiate default collector
  c := colly.NewCollector(/*colly.AllowedDomains("9to5mac.com", "9to5mac.com"),*/)

  // On every a element which has selector attribute scrape the link and title
  c.OnHTML(selector, func(e *colly.HTMLElement) {
    link := e.Attr("href")
    title := e.Text
    
    // Add article to pagelinks slice
    article := PageLink{id:id, link:link, title: title}
    articles = append(articles, article)
    id++
  })

  // Start scraping on tarURL
  err := c.Visit(targetUrl)

  // Scraper error print
  if err != nil {
    fmt.Println("Scraping Error:", err)
  }

  // Print scraper statues
  fmt.Printf("\nc: %v\n", c)
  return articles 
}

