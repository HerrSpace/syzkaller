// Copyright 2021 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/syzkaller/pkg/compiler"
)

type serenity struct{}

func (*serenity) prepare(sourcedir string, build bool, arches []*Arch) error {
	return nil
}

func (*serenity) prepareArch(arch *Arch) error {
	return nil
}

func (*serenity) processFile(arch *Arch, info *compiler.ConstInfo) (map[string]uint64, map[string]bool, error) {
	args := []string{
		"-I", filepath.Join(arch.sourceDir, "..", "Build", "i686", "Root", "usr", "include"),
		"-I", filepath.Join(arch.sourceDir, "..", "Userland", "Libraries", "LibC"),
	}
	for _, incdir := range info.Incdirs {
		args = append(args, "-I"+filepath.Join(arch.sourceDir, incdir))
	}
	fmt.Printf("dirs: %v", arch.includeDirs)
	if arch.includeDirs != "" {
		for _, dir := range strings.Split(arch.includeDirs, ",") {
			args = append(args, "-I"+dir)
		}
	}
	params := &extractParams{
		TargetEndian:   arch.target.HostEndian,
		ExtractFromELF: true,
	}
	cc := arch.target.CCompiler
	return extract(info, cc, args, params)
}
