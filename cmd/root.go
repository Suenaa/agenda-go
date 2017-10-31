package cmd

import (
	"fmt"
	"os"
	"log"

	"github.com/spf13/cobra"
	logs "github.com/Suenaa/agenda-go/logs"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "agd",
	Short: "Agenda-go",
	Long:  "A tool for managing meetings",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		logs.Log(err)
		os.Exit(1)
	}
}
