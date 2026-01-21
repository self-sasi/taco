package state

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/self-sasi/taco/internal/utils"
)

const lockFileName = "taco.push.lock"

func WriteLock() error {
	gitDir, err := utils.GitDir()
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(gitDir, lockFileName), []byte("locked\n"), 0o644)
}

func RemoveLock() error {
	gitDir, err := utils.GitDir()
	if err != nil {
		return err
	}
	p := filepath.Join(gitDir, lockFileName)
	if err := os.Remove(p); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func Status() (bool, error) {
	gitDir, err := utils.GitDir()
	if err != nil {
		return false, err
	}

	p := filepath.Join(gitDir, lockFileName)

	_, err = os.Stat(p)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}
