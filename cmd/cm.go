package cmd

import (
	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "create a meeting",
	Long:  `create a meeting`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(cmCmd)

	cmCmd.Flags().StringP("title", "t", "", "title of the meeting")
	cmCmd.Flags().StringSliceP("participant", "p", nil, "participants of the meeting")
	cmCmd.Flags().StringP("start", "s", "", "when to start the meeting")
	cmCmd.Flags().StringP("end", "e", "", "when to end the meeting")
}
