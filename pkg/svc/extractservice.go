package svc

import (
	"os"
	_ "github.com/pojntfx/k3sd/pkg/svc/statik" // Embedded k3s binary
	"github.com/rakyll/statik/fs"
)

// ExtractService is a service that extracts binaries.
type ExtractService struct {
	BinaryDir          string
	BinaryInternalPath string
}

// Extract extracts the binary.
func (e *ExtractService) Extract() error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	data, err := fs.ReadFile(statikFS, e.BinaryInternalPath)
	if err != nil {
		return err
	}

	binaryFile, err := os.Create(e.BinaryDir)
	if err != nil {
		return err
	}

	if _, err = binaryFile.Write(data); err != nil {
		return err
	}
	defer binaryFile.Close()

	if err := os.Chmod(e.BinaryDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// Cleanup deletes the extracted binary.
func (e *ExtractService) Cleanup() error {
	return os.Remove(e.BinaryDir)
}
