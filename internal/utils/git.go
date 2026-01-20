package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// RunGit is a utility method for running `git` with the provided flags
// or subcommands. It returns a ("", error) tuple if the command executuon
// fails, (output, nil) if successful.
func runGit(args ...string) (string, error) {
	gitCmd := exec.Command("git", args...)

	output, err := gitCmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// VerifyInsideGitWorkTree runs `git rev-parse --is-inside-work-tree`
// to check that we are in a git initialized repo. Returns error if not.
func VerifyInsideGitWorkTree() error {
	output, err := runGit("rev-parse", "--is-inside-work-tree")
	if err != nil {
		return err
	}

	if output == "false" {
		return fmt.Errorf(`Not a Git Repo`)
	}
	return nil
}

// GitDir returns the repo's git dir path (e.g. ".git" or a worktree git dir).
func GitDir() (string, error) {
	return runGit("rev-parse", "--git-dir")
}
