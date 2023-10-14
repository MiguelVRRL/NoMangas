package utils

import (
	"fmt"
	"os/exec"
)


func List () (error) {
  cmd := exec.Command("ls", Home)
  out,err := cmd.Output()
  if err != nil {
    return err
  }
  fmt.Println(string(out))
  return nil
}
