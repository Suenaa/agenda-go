package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lsuCmd represents the lsu command
var lsuCmd = &cobra.Command{
	Use:   "lsu",
	Short: "list all users",
	Long:  `list all users`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lsu called")
	},
}

func init() {
	RootCmd.AddCommand(lsuCmd)
}
