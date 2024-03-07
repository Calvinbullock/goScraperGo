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
  var scrapeFlag string
  fmt.Println("Do you want to scrape? (y/n): ")
  fmt.Scanf("%f", scrapeFlag)

  db := connectDataBase()

  if scrapeFlag == "y"{
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
func connectDataBase() *sql.DB{
  // NOTE Database connection details - Generaly this should not be hard codded.
  //   pass this in insted of hard codeing
  dsn := "debian-sys-maint:fAkBoMzWqEDlkGNf@tcp(localhost:3306)/go_scrape?charset=utf8mb4"

  // Connect to the database
  db, err := sql.Open("mysql", dsn)
  
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
  id := 0
  // Instantiate default collector
  c := colly.NewCollector(
    //colly.AllowedDomains("9to5mac.com", "9to5mac.com"),
  )
 
  var articles []PageLink

  // On every a element which has selector attribute scrape the link and title
  c.OnHTML(selector, func(e *colly.HTMLElement) {
    link := e.Attr("href")
    title := e.Text
    
    page := PageLink{id:id, link:link, title: title}
    articles = append(articles, page)
    id++
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

