package utils

import (
   "log"
   "github.com/gocolly/colly"
)

func Search() error {
  c := colly.NewCollector()
  log.Println("Manga searched")
  if err := Download(c); err != nil  {
    return err
  }
  log.Printf("manga Downloaded at %s\n", "dir")
  return nil
}

func Download(c *colly.Collector) error {
  return nil 
}
