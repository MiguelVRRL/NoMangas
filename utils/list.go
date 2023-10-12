package utils

import "os/exec"


func List () (error) {
  cmd := exec.Command("ls", Dir)
  return cmd.Run()
}
