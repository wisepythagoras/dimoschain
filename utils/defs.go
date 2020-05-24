package utils

import (
	"os"
)

// GetChainDir returns the directory the blockchain should live in.
func GetChainDir(createIfNotExists bool) (string, error) {
	// Get the home directory from the os package.
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	path := home + "/.dimos"

	if createIfNotExists {
		// Create the directory if it doesn't exist.
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModeDir)
		}
	}

	return path, nil
}
