package cli

import "github.com/spf13/cobra"

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status shows the state of the state of the current repo (lonked/unlocked)",
	Long:  "placeholder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
