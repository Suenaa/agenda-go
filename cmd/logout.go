package cmd

import (
	"fmt"

	"github.com/Andiedie/agenda-go/service"
	"github.com/Andiedie/agenda-go/tools"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out",
	Long:  `log out agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.UserLogout(); err == nil {
			fmt.Println("Success")
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}
