// Copyright 2022 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package vcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClangVersion(t *testing.T) {
	binDir := "/some/dir/"
	tags := make(map[string]bool)

	// No tags case.
	actual := linuxClangDir(tags, binDir)
	expected := binDir + "llvm-9.0.1/bin"
	assert.Equal(t, actual, expected, "unexpected clang dir")

	// Recent tag case.
	tags["v5.9"] = true
	actual = linuxClangDir(tags, binDir)
	expected = ""
	assert.Equal(t, actual, expected, "unexpected clang dir")
}

func TestGCCVersion(t *testing.T) {
	binDir := "/some/dir/"
	tags := make(map[string]bool)

	// No tags case.
	actual := linuxGCCDir(tags, binDir)
	expected := binDir + "gcc-5.5.0/bin"
	assert.Equal(t, actual, expected, "unexpected gcc dir")

	// Somewhat old tag case.
	tags["v4.12"] = true
	actual = linuxGCCDir(tags, binDir)
	expected = binDir + "gcc-8.1.0/bin"
	assert.Equal(t, actual, expected, "unexpected gcc dir")

	// Recent tag case.
	tags["v5.16"] = true
	actual = linuxGCCDir(tags, binDir)
	expected = ""
	assert.Equal(t, actual, expected, "unexpected gcc dir")
}
