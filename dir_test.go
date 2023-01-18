// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/golistic/xt"
)

func TestIsDir(t *testing.T) {
	t.Run("existing directory", func(t *testing.T) {
		xt.Assert(t, IsDir("_testdata"))
	})

	t.Run("non-existing directory", func(t *testing.T) {
		xt.Assert(t, !IsDir("nonexistingdirectory"))
	})

	t.Run("regular file is not a dir", func(t *testing.T) {
		xt.Assert(t, !IsDir("dir.go"))
	})
}

func TestRegularFilesInDir(t *testing.T) {
	t.Run("list regular files", func(t *testing.T) {
		dir := "_testdata/regular_files_in_dir"
		readdir, err := os.ReadDir(dir)
		xt.OK(t, err)
		var exp []string
		for _, f := range readdir {
			if !f.Type().IsRegular() {
				continue
			}
			exp = append(exp, f.Name())
		}

		files, err := RegularFilesInDir(dir)
		xt.OK(t, err)

		xt.Eq(t, exp, files)
	})
}

func TestRegularFilesInDirWithFullPath(t *testing.T) {
	t.Run("list regular files with full path", func(t *testing.T) {
		dir := "_testdata/regular_files_in_dir"
		readdir, err := os.ReadDir(dir)
		xt.OK(t, err)
		var exp []string
		for _, f := range readdir {
			if !f.Type().IsRegular() {
				continue
			}
			exp = append(exp, filepath.Join(dir, f.Name()))
		}

		files, err := RegularFilesInDirWithFullPath(dir)
		xt.OK(t, err)

		xt.Eq(t, exp, files)
	})
}
