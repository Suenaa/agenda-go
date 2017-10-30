package cmd

import (
	"github.com/spf13/cobra"
)

// registCmd represents the regist command
var registCmd = &cobra.Command{
	Use:   "regist",
	Short: "regist a new user",
	Long:  `regist a new user`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(registCmd)

	registCmd.Flags().StringP("username", "u", "", "the username you want")
	registCmd.Flags().StringP("password", "p", "", "the password you want")
	registCmd.Flags().StringP("email", "e", "", "your email address")
	registCmd.Flags().StringP("telephone", "t", "", "your telephone number")
}
