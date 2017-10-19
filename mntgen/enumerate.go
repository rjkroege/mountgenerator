package mntgen

import (
	"path/filepath"
)

// EnumerateDisks finds the disks and enumerates them.
func EnumerateDisks(path string) ([]string, error) {

	globpath := filepath.Join(path, "google-*")
	disknames, err := filepath.Glob(globpath)
	if err != nil {
		return []string{}, err
	}

	filtered := []string{}
	for _, f := range disknames {
		m, err := filepath.Match(filepath.Join(path, "google-persistent-disk*"), f)
		if err != nil {
			return []string{}, err
		}
		if !m {
			filtered = append(filtered, f)
		}
	}

	return filtered, nil

}
