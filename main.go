
package main

import (
  "fmt"

  "github.com/gocolly/colly"
)

func main() {
  // Your program code goes here
  c := colly.NewCollector()
  fmt.Println(c)
  fmt.Println("working")
}

