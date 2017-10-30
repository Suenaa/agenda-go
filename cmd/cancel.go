package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel a meeting",
	Long:  `cancel a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cancel called")
	},
}

func init() {
	RootCmd.AddCommand(cancelCmd)

	cancelCmd.Flags().StringP("title", "t", "", "title of the meeting")
}
