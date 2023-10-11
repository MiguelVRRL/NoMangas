package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/signintech/gopdf"
)
var rect = gopdf.Rect{H:3371,W:1000}
var cap_page string = "https://leermanga.net/capitulo/"
var menu_page string = "https://leermanga.net/manga/"

func Search(args []string) error {
  name, chapters := args[0],args[1:]
  if areNumbers := AreNumbers(chapters); areNumbers != nil {
    return areNumbers
  }
  c := colly.NewCollector()
  log.Println("Manga Searched")
  name = strings.ReplaceAll(name, " ", "-")
  for i := 0; i < len(chapters); i++ {
    if err := Download(c,name,chapters[i]); err != nil {
      return err
    }
  }
  log.Printf("manga Downloaded at %s\n", "dir")

  return nil
}


func Download(collyC *colly.Collector,name string, chapter string) error  {
  var imageListChapter []string
  var manwha string
  c := colly.NewCollector()
  collyC.OnHTML(".lazyload", func(e *colly.HTMLElement) { 
    listLinks := strings.Split(e.Attr("data-src"),"\n")
    imageListChapter = append(imageListChapter, listLinks...)
  })
  collyC.Visit(cap_page+name+"-"+chapter+".00")
  c.OnHTML(".manga-title-info",func(e *colly.HTMLElement) { 
    if e.Text == "Manwha" {
      manwha = e.Text
    }
  }) 
  c.Visit(menu_page+name)
  if len(imageListChapter) == 0 {
    return errors.New("Manga not found.")
  }
  if manwha == "Manhwa" {
    if err := ImageToPDF(name,chapter,imageListChapter,gopdf.Rect{H:5600.00,W:100.00}); err != nil {
      return err   
    } 
  } else {
    if err := ImageToPDF(name,chapter,imageListChapter,*gopdf.PageSizeLetter); err != nil {
      return err   
    }
  }
  return nil
}


func ImageToPDF(name,chapter string ,links []string, rect gopdf.Rect) error {
  pdf := gopdf.GoPdf{}
  pdf.Start(gopdf.Config{PageSize: rect})  
  x :=1.00
  y :=1.00
  pdf.AddPage()
  for i := 0; i < len(links); i++ {
    imageBytes,err := DownloadFile(links[i])
    if err != nil {
        return err
    }
    b, err := gopdf.ImageHolderByBytes(imageBytes)
    if err != nil {
      return err
    }
    pdf.ImageByHolder(b, x, y, &rect)
    pdf.AddPage()
  }
  

	pdf.WritePdf("./"+name+"-capitulo-"+chapter+".00"+".pdf")
  
  return nil
}
func DownloadFile( url string) ([]byte,error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil,err
  }
  defer resp.Body.Close()
  b,_ := io.ReadAll(resp.Body)
  if err != nil {
    return nil,err
  }
  return b,nil

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


