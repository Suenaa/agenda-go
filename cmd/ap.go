package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Andiedie/agenda-go/service"
	"github.com/Andiedie/agenda-go/tools"
	"github.com/spf13/cobra"
)

// apCmd represents the ap command
var apCmd = &cobra.Command{
	Use:   "ap",
	Short: "add participants to a meeting",
	Long:  `add participants to a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participants, _ := cmd.Flags().GetStringSlice("participant")
		if title == "" {
			tools.Report(errors.New("title required"))
		}
		if participants == nil || len(participants) == 0 {
			tools.Report(errors.New("participant(s) required"))
		}
		noError := true
		for _, one := range participants {
			if err := service.AddParticipator(title, one); err != nil {
				noError = false
				fmt.Fprintln(os.Stderr, err)
			}
		}
		if noError {
			fmt.Println("Success")
		}
	},
}

func init() {
	RootCmd.AddCommand(apCmd)

	apCmd.Flags().StringP("title", "t", "", "title of the meeting")
	apCmd.Flags().StringSliceP("participant", "p", nil, "participants want to add")
}
