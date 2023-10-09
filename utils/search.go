package utils

import (
	"errors"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

var web_page string = "https://my.ninemanga.com/manga/"

func Search(args []string) error {
  name, chapters := args[0],args[1:]
  if areNumbers := AreNumbers(chapters); areNumbers != nil {
    return areNumbers
  }
  c := colly.NewCollector()
  log.Println("Manga Searched")
  if err := Download(c,name); err != nil  {
    return err
  }
  log.Printf("manga Downloaded at %s\n", "dir")
  name = strings.ReplaceAll(name, " ", "+")
  c.Visit(web_page+name+".html")
  return nil
}

func AreNumbers(arg []string) error {
  for i,s := range arg { 
    for _,r := range s {
      if r < 48 || r > 57 {
        return errors.New("introduce the num of the chapter that you wish to download. Args malformed " + arg[i] )    
      }
    }
  } 
  return nil
}


func Download(c *colly.Collector, name string) error {
  
  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    l:= strings.Split(e.Text, "\n")
    log.Println(l)
	})
  return nil 
}


func ImageToPDF() error {
  return nil
}
