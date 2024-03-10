# Overview

This program will scrape a url and then store all the scraped data into a database. 

My goal in wrting this program was to learn about scraping and database interfaces. As well as to help me remember how to use mySQL
and to get to know mySQL server.

[Software Demo Video](https://youtu.be/pxG08mXEJXc)

# Development Environment

Language: go-lang, mySQL
Library: colly, go-sql-driver/mysql/
Editers: NeoVim, mySQL-server

# Useful Websites

### videos
- [A web scraper video totorial with go.] (https://pkg.go.dev/github.com/gocolly/colly#section-readme)
- [Intro to wroking with mySQL server] (https://www.youtube.com/watch?v=xiUTqnI6xk8)
- [Working with sql in go] (https://www.youtube.com/watch?v=Y7a0sNKdoQk)
- [Basics of web scarping] (https://www.youtube.com/watch?v=EJNJ5_i_zu8)
)
### Docs
- [Colly Docs - scaper libray] (https://www.youtube.com/watch?v=LMPeAttF2ng&list=PL5dTjWUk_cPbbCYRQKhPnmougbStBPba8&index=6)
- [Go-docs building a module] (https://go.dev/doc/tutorial/create-module)
- [Go-docs server drivers] (https://pkg.go.dev/database/sql)

# Future Work

- Data Base search
- Make the scraper more user configuralble
    - prompt for target url
    - prompt for target html element
