package utils

import (
  "io"
	"errors"
	"log"
	"strings"
  "net/http"

  "github.com/signintech/gopdf"
	"github.com/gocolly/colly"
)

var web_page string = "https://leermanga.net/capitulo/"

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
  collyC.OnHTML(".lazyload", func(e *colly.HTMLElement) { 
    listLinks := strings.Split(e.Attr("data-src"),"\n")
    imageListChapter = append(imageListChapter, listLinks...)
  })
  collyC.Visit(web_page+name+"-"+chapter+".00")
  if len(imageListChapter) == 0 {
    return errors.New("Manga not found.")
  }
  if err := ImageToPDF(name,chapter,imageListChapter); err != nil {
    return err   
  }
  return nil
}


func ImageToPDF(name,chapter string ,links []string) error {
  pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4 })  
  x :=1.00
  y :=1.00
  for i := 0; i < len(links); i++ {
    imageBytes,err := DownloadFile(links[i])
    if err != nil {
        return err
    }
    b, err := gopdf.ImageHolderByBytes(imageBytes)
    if err != nil {
      return err
    }
    pdf.ImageByHolder(b,x,y,nil)
    y += 516 
  }
  
	pdf.SetXY(250, 200) //move current location

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


