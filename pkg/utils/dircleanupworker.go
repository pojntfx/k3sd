package utils

import "os"

// DirCleanupWorker is a worker that cleans up directories.
type DirCleanupWorker struct {
	DirsToClean []string
}

// CleanupDirs cleans up the directories.
func (d *DirCleanupWorker) CleanupDirs() error {
	for _, dir := range d.DirsToClean {
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}

	return nil
}
