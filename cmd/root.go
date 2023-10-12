package cmd

import (
	"log"
	"os"

	"github.com/MiguelVRRL/nomangas/utils"
	"github.com/spf13/cobra"
)




var rootCmd = &cobra.Command{
	Use:   "nomangas",
	Short: "A brief description of your application",
  Long: ``,

	Run: func(_ *cobra.Command, _ []string) {
    if err := utils.List(); err != nil {
      log.Fatal(err.Error())
    }
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


