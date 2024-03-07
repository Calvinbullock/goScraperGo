package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

type PageLink struct {
  id int
  title string
  link string
}

func main() {
  // User input for scraping or not.
  var userChoice int

  fmt.Println("Do you want to scrape url  (1): ")
  fmt.Println("Print out scraped data     (2): ")
  fmt.Println("Exit                       (3): ")
  fmt.Scanf("%n\n", userChoice)

  // NOTE Database connection details - Generaly this should not be hard codded.
  dbCredentals := "debian-sys-maint:fAkBoMzWqEDlkGNf@tcp(localhost:3306)/go_scrape?charset=utf8mb4"
  tarURL := "https://9to5mac.com"
  db := connectDataBase(dbCredentals)
  var articles []PageLink

  // Scrape the url and store into DataBase
  if userChoice == 1 {
    // send scarper to tarURL.
    articles := scrapeUrl(tarURL, "a.article__title-link")

    // send scraped results to DataBase.
    for _, article := range articles {
      insertToDatabase(db, article)
    }
  }

  // TODO need to make a function to re-parse the database when realoading this program.
  if userChoice == 2 {
    printArticles(articles)
  }
  
  // Close the DataBase connection and exit
  if userChoice == 3 { 
    db.Close() 
    return
  }
}

// Print out scraped Data
func printArticles(articles []PageLink) {
  if articles != nil {
    fmt.Println("Collected links:")
    for _, article := range articles {
      fmt.Printf("\n%d \n%s \n%s\n", article.id, article.title, article.link)
    }
  }
}

// TODO set up a database search
func linkSearch(articles []PageLink, searchTarget string) {
  
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

