package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all meeting you create",
	Long:  `clear all meeting you create`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clear called")
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
