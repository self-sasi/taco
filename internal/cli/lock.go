package cli

import "github.com/spf13/cobra"

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "A brief description of your application",
	Long:  `lock locks your repo from pushing to remote`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
