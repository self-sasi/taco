package cli

import (
	"fmt"

	"github.com/self-sasi/taco/internal/state"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "A brief description of your application",
	Long:  `lock locks your repo from pushing to remote`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := state.InstallPrePush(); err != nil {
			fmt.Printf(`installation error: %v`, err)
			return
		}
		if err := state.WriteLock(); err != nil {
			fmt.Printf(`lock installation error: %v`, err)
			return
		}
		fmt.Printf(`locked the repo`)
	},
}

func init() {
	rootCmd.AddCommand(lockCmd)
}
