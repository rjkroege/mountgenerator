package mntgen

import (
	"path/filepath"
)


// EnumerateDisks finds the disks and enumerates them.
func EnumerateDisks(path string) ([]string, error) {

	globpath  := filepath.Join(path, "google-*")
	disknames, err := filepath.Glob(globpath)
	if err != nil {
		return []string{}, err
	}

	// TODO(rjk): insert filtering
	return disknames, nil

}

