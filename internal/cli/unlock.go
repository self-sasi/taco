package cli

import (
	"github.com/self-sasi/taco/internal/utils"
	"github.com/spf13/cobra"
)

var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "unlock unlocks the git repository and allows you to push to remote",
	Long:  "placeholder",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := utils.VerifyInsideGitWorkTree(); err != nil {
			return
		}
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(unlockCmd)
}
