package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apCmd represents the ap command
var apCmd = &cobra.Command{
	Use:   "ap",
	Short: "add participants to a meeting",
	Long:  `add participants to a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ap called")
	},
}

func init() {
	RootCmd.AddCommand(apCmd)

	apCmd.Flags().StringP("title", "t", "", "title of the meeting")
	apCmd.Flags().StringSliceP("participant", "p", nil, "participants want to add")
}
