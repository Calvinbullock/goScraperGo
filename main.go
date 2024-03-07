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
  scrapeFlag := true

  db := connectDataBase()

  // NOTE keep this from runnint for now
  if scrapeFlag {
    // send scarper to tarURL.
    tarURL := "https://9to5mac.com"
    articles := scrapeUrl(tarURL, "a.article__title-link")

    // print scraped results.
    fmt.Println("Collected links:")
    for _, article := range articles {
      fmt.Printf("\n%d \n%s \n%s\n", article.id, article.title, article.link)
      insertToDatabase(db, article)
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
func insertToDatabase(db *sql.DB, article PageLink) error {
/* NOTE This function was writtn with the help of google Bard.
//  This is the prompt I used:
//    Can you show me an exsample of go-lang code that is inserting data into a 
//    my sql data base? */

  // Prepare the SQL statement with placeholders for values
  stmt, err := db.Prepare("INSERT INTO article_table(id, title, link) VALUES ($1, $2, $3)")
  if err != nil {
    return err
  }
  defer stmt.Close() // Close the prepared statement

  // Execute the statement with the articles's data
  result, err := stmt.Exec(article.id, article.title, article.link)
  if err != nil {
    log.Fatal(err)
    return err
  }

  // Get the last inserted ID (optional)
  lastID, err := result.LastInsertId()
  if err != nil {
    return err
  }
  fmt.Printf("Last inserted ID: %d\n", lastID)

  return nil
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

