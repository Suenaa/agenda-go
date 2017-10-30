package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lsmCmd represents the lsm command
var lsmCmd = &cobra.Command{
	Use:   "lsm",
	Short: "list all meetings during a period",
	Long:  `list all meetings during a period`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lsm called")
	},
}

func init() {
	RootCmd.AddCommand(lsmCmd)

	lsmCmd.Flags().StringP("start", "s", "", "start time")
	lsmCmd.Flags().StringP("end", "e", "", "end time")
}
