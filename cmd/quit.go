package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// quitCmd represents the quit command
var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "quit a meeting",
	Long:  `quit a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quit called")
	},
}

func init() {
	RootCmd.AddCommand(quitCmd)

	quitCmd.Flags().StringP("title", "t", "", "title of the meeting")
}
