// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"bytes"
	"os"
)

func linuxControlGroup() []byte {
	data, _ := os.ReadFile("/proc/1/cgroup") // error skipped; we use zero value
	return data
}

// InDockerContainer returns true if the application is running within a Docker container.
func InDockerContainer() bool {
	return bytes.Contains(linuxControlGroup(), []byte(":/docker/"))
}
