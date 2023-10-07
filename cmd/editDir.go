package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var editDirCmd = &cobra.Command{
	Use:   "editDir",
	Short: "A brief description of your command",
  Long: ``,
  SuggestionsMinimumDistance: 3,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("editDir called")
	},
}

func init() {
	rootCmd.AddCommand(editDirCmd)

}
