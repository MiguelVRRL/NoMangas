package utils

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

func ReadMangas(name, chapter string) error {
  name = strings.ReplaceAll(name," ","-")
  manga := Home+name+"-capitulo-"+chapter+".00.pdf"
  log.Println(manga)
  c := exec.Command("xdg-open", manga )
  if c.Run() != nil {
    return errors.New("We can't open the file")
  }
  return nil
}
