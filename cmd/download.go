package cmd

import (
	//	"fmt"

	"log"

	"github.com/MiguelVRRL/nomangas/utils"
	"github.com/spf13/cobra"
)


var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "the command to download the chapter of anything manga",
  Args: cobra.MinimumNArgs(2),
  Long: ``,
  SuggestionsMinimumDistance: 3,
	Run: func(cmd *cobra.Command, args []string) {
    if err := utils.Search(); err != nil {
      log.Printf("Error: %s", err.Error())
    }
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
