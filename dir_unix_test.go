// Copyright (c) 2020, Geert JM Vanderkelen

package xos

import (
	"os"
	"os/exec"
	"testing"

	"github.com/golistic/xt"
)

func TestFilesInDir(t *testing.T) {
	t.Run("all files with symbolic link", func(t *testing.T) {
		dir := "_testdata/files_in_dir"

		script := dir + "/create_sym_link.sh"
		xt.OK(t, os.Chmod(script, 0700))
		cmd := exec.Command(dir + "/create_sym_link.sh")
		xt.OK(t, cmd.Run())

		readdir, err := os.ReadDir(dir)
		xt.OK(t, err)
		var exp []string
		for _, f := range readdir {
			exp = append(exp, f.Name())
		}

		files, err := FilesInDir(dir)
		xt.OK(t, err)

		xt.Eq(t, exp, files)
	})
}
