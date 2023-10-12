package utils

import (
	"os/exec"
	"strings"
)

func ReadMangas(name, chapter string) error {
  name = strings.ReplaceAll(name," ","-")
  manga := name+"-Capitulo-"+chapter+".pdf"
  exec.Command("xdg-open", manga )
  return nil
}
