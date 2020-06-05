package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

// CheckIfFileExists checks if the given path exists.
func CheckIfFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// GetChainDir returns the directory the blockchain should live in.
func GetChainDir(createIfNotExists bool) (string, error) {
	// Get the home directory from the os package.
	home, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	path := home + "/" + DimosDir

	if createIfNotExists {
		// Create the directory if it doesn't exist.
		if !CheckIfFileExists(path) {
			os.Mkdir(path, 0777)
		}
	}

	return path, nil
}

// ReadFileInChainDir Read a file in the blockchain directory.
func ReadFileInChainDir(fileName string) ([]byte, error) {
	// Get the chain's directory.
	home, err := GetChainDir(true)

	if err != nil {
		return nil, err
	}

	// Compose the target file's path.
	path := home + "/" + ChainDir + "/" + fileName

	// Check if it exists.
	if !CheckIfFileExists(path) {
		return nil, errors.New("The blockchain has not been instanciated yet")
	}

	// Return the contents.
	return ioutil.ReadFile(path)
}

// GetGenesisHash gets the genesis block hash from inside
func GetGenesisHash() ([]byte, error) {
	return ReadFileInChainDir(Genesis)
}

// GetCurrentHash gets the current block's hash in the blockchain.
func GetCurrentHash() ([]byte, error) {
	return ReadFileInChainDir(CurrentHash)
}

// WriteCurrentHash writes the current hash to the disk.
func WriteCurrentHash(hash []byte) error {
	// Get the chain's directory.
	home, err := GetChainDir(true)

	if err != nil {
		return err
	}

	// Compose the target file's path.
	path := home + "/" + ChainDir + "/" + CurrentHash

	// Write the hash.
	return WriteToFile(path, hash)
}

// WriteGenesisHash writes the genesis hash to the disk.
func WriteGenesisHash(hash []byte) error {
	// Get the chain's directory.
	home, err := GetChainDir(true)

	if err != nil {
		return err
	}

	// Compose the target file's path.
	path := home + "/" + ChainDir + "/" + Genesis

	// Write the hash.
	return WriteReadOnlyFile(path, hash)
}

// WriteToFile writes the contents to the file at the given path.
func WriteToFile(path string, contents []byte) error {
	return ioutil.WriteFile(path, contents, 0644)
}

// ReadFile reads the contents of the file at the given path.
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// WriteReadOnlyFile writes the contents to the read only file at the given path.
func WriteReadOnlyFile(path string, contents []byte) error {
	return ioutil.WriteFile(path, contents, 0444)
}
