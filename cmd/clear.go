package cmd

import (
	"fmt"

	"github.com/Andiedie/agenda-go/service"
	"github.com/Andiedie/agenda-go/tools"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all meeting you create",
	Long:  `clear all meeting you create`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.DeleteAllMeeting(); err == nil {
			fmt.Println("Success")
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
