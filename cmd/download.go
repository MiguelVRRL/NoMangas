package cmd

import (
	"fmt"
	"log"

	"github.com/MiguelVRRL/nomangas/utils"
	"github.com/spf13/cobra"
)

var limit bool

var downloadCmd = &cobra.Command{
	Use:   "download [name] [num of each chapter to download]",
	Short: "the command to download the chapter of anything manga",
  Long: ``,
  Args: func(_ *cobra.Command, args []string) error {
    if limit {
      return nil
    }
    if len(args) < 2 {
      return fmt.Errorf("Introduce minimum 2 args; actual args: %d", len(args))
    }
    if len(args) > 100 {
      return fmt.Errorf("Maximum of chapter to download is 10")
    }
    return nil
  },

  SuggestionsMinimumDistance: 3,
	Run: func(_ *cobra.Command, args []string) {
    if err := utils.Search(args); err != nil {
      log.Printf("Error: %s", err.Error())
    } else {
    log.Println("Manga Downloaded")
    }
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
  downloadCmd.PersistentFlags().BoolVarP(&limit, "limit","l", false, "flag to remove the limit of download")
}
