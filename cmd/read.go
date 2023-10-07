package cmd

import (

  "github.com/MiguelVRRL/nomangas/utils"
	"github.com/spf13/cobra"
)


var readCmd = &cobra.Command{
	Use:   "read",
	Short: "A brief description of your command",
  Long: ``,
  SuggestionsMinimumDistance: 3,

	Run: func(cmd *cobra.Command, args []string) {
		utils.ReadMangas()
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
