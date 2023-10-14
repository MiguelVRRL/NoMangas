package utils

import "os/exec"


func List () (error) {
  cmd := exec.Command("ls", Home)
  return cmd.Run()
}
