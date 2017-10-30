package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dpCmd represents the dp command
var dpCmd = &cobra.Command{
	Use:   "dp",
	Short: "delete participants in a meeting",
	Long:  `delete participants in a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dp called")
	},
}

func init() {
	RootCmd.AddCommand(dpCmd)

	dpCmd.Flags().StringP("title", "t", "", "title of the meeting")
	dpCmd.Flags().StringSliceP("participant", "p", nil, "participants want to remove")
}
