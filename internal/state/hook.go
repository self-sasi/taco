package state

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/self-sasi/taco/internal/utils"
)

const hookSentinel = "managed-by: taco"

//go:embed scripts/pre-push.sh
var prePushHook string

func InstallPrePush() error {
	err := utils.VerifyInsideGitWorkTree()
	if err != nil {
		return err
	}

	gitDir, err := utils.GitDir()
	if err != nil {
		return err
	}

	hooksDir := filepath.Join(gitDir, "hooks")
	hookPath := filepath.Join(hooksDir, "pre-push")

	if err := os.MkdirAll(hooksDir, 0o755); err != nil {
		return fmt.Errorf("create hooks dir: %w", err)
	}

	if b, err := os.ReadFile(hookPath); err == nil {
		if strings.Contains(string(b), hookSentinel) {
			return nil
		}
		return fmt.Errorf("pre-push already exists and is not managed by taco: %s", hookPath)
	}

	if err := os.WriteFile(hookPath, []byte(prePushHook), 0o755); err != nil {
		return fmt.Errorf("write pre-push: %w", err)
	}

	if runtime.GOOS != "windows" {
		_ = os.Chmod(hookPath, 0o755)
	}
	return nil
}
