// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"bytes"
	"testing"

	"github.com/golistic/xt"
)

func TestInDockerContainer(t *testing.T) {
	// this works differently when in container or not
	exp := bytes.Contains(linuxControlGroup(), []byte(":/docker/"))
	xt.Eq(t, exp, InDockerContainer())
}
