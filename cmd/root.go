
package cmd

import (
	"os"
  "fmt"

  "github.com/MiguelVRRL/nomangas/utils"
	"github.com/spf13/cobra"
)




var rootCmd = &cobra.Command{
	Use:   "nomangas",
	Short: "A brief description of your application",
  Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(utils.List)
  },
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


