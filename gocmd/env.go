package gocmd

import (
	"os"
	"path/filepath"
)

const (
	EnvNameGoPath = "GOPATH"
	EnvNameGoBin  = "GOBIN"
)

const (
	localBinDirectory = "bin"
	goBinDirectory    = "bin"
)

// SetLocalGoBin sets the $GOBIN variable to `$(pwd)/bin`.
func SetLocalGoBin() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	localBinPath := filepath.Join(cwd, localBinDirectory)
	return os.Setenv(EnvNameGoBin, localBinPath)
}

// SetCustomGoBin sets the $GOBIN to the provided path
func SetCustomGoBin(path string) error {
	cleanedPath := filepath.Clean(path)
	return os.Setenv(EnvNameGoBin, cleanedPath)
}

// ResetGoBin resets the $GOBIN to the default location `$GOPATH/bin`
func ResetGoBin() error {
	defaultGoBin := filepath.Join(os.Getenv(EnvNameGoPath), goBinDirectory)
	return os.Setenv(EnvNameGoBin, defaultGoBin)
}
