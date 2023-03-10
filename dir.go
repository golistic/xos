// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"os"
	"path/filepath"
	"sort"
)

// IsDir returns whether path is a directory.
func IsDir(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		return fi.Mode().IsDir()
	}
	return false
}

// RegularFilesInDir returns regular files found in directory path.
func RegularFilesInDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var l []string
	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}
		l = append(l, entry.Name())
	}

	sort.Strings(l)
	return l, nil
}

// RegularFilesInDirWithFullPath returns regular files found in directory path with
// each path included in the filename.
func RegularFilesInDirWithFullPath(path string) ([]string, error) {
	files, err := RegularFilesInDir(path)
	if err != nil {
		return nil, err
	}

	absFiles := make([]string, len(files))
	for i, f := range files {
		absFiles[i] = filepath.Join(path, f)
	}
	return absFiles, nil
}

// FilesInDir returns file which are not directories found in directory path.
func FilesInDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var l []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		l = append(l, entry.Name())
	}

	sort.Strings(l)
	return l, nil
}
