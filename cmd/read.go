package cmd

import (
	"log"

	"github.com/MiguelVRRL/nomangas/utils"
	"github.com/spf13/cobra"
)


var readCmd = &cobra.Command{
	Use:   "read",
	Short: "A brief description of your command",
  Long: ``,
  SuggestionsMinimumDistance: 3,
  Args: cobra.MaximumNArgs(2),
	Run: func(_ *cobra.Command, args []string) {
    if err := utils.ReadMangas(args[0],args[1]); err != nil {
      log.Fatal(err.Error())
    }
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
