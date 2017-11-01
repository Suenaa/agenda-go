package cmd

import (
	"fmt"

	"github.com/Andiedie/agenda-go/service"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-go/logs"
)

// lsuCmd represents the lsu command
var lsuCmd = &cobra.Command{
	Use:   "lsu",
	Short: "list all users",
	Long:  `list all users`,
	Run: func(cmd *cobra.Command, args []string) {
		users := service.QueryAllUsers()
		for u := range users {
			fmt.Println(u)
		}
		logs.EventLog("list all users")
	},
}

func init() {
	RootCmd.AddCommand(lsuCmd)
}
