package cli

import (
	"fmt"

	"github.com/self-sasi/taco/internal/state"
	"github.com/spf13/cobra"
)

var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "unlock unlocks the git repository and allows you to push to remote",
	Long:  "placeholder",
	Run: func(cmd *cobra.Command, args []string) {
		if err := state.InstallPrePush(); err != nil {
			fmt.Printf(`installation error: %v`, err)
			return
		}
		if err := state.RemoveLock(); err != nil {
			fmt.Printf(`lock removal error: %v`, err)
			return
		}
		fmt.Printf(`unlocked the repo`)
	},
}

func init() {
	rootCmd.AddCommand(unlockCmd)
}
