package cli

import (
	"fmt"

	"github.com/self-sasi/taco/internal/state"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status shows the state of the state of the current repo (lonked/unlocked)",
	Long:  "placeholder",
	Run: func(cmd *cobra.Command, args []string) {
		if err := state.InstallPrePush(); err != nil {
			fmt.Printf(`installation error: %v`, err)
			return
		}

		isLocked, err := state.Status()
		if err != nil {
			fmt.Printf(`status retrieval error: %v`, err)
			return
		}

		if isLocked {
			fmt.Printf(`repo is locked`)
		} else {
			fmt.Printf(`repo is unlocked`)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
