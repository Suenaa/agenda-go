package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete current account",
	Long:  `delete current account`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("del called")
	},
}

func init() {
	RootCmd.AddCommand(delCmd)

	delCmd.Flags().StringP("password", "p", "", "your password")
}
