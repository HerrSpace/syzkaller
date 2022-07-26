// Copyright 2019 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package vcs

import (
	"fmt"

	"github.com/google/syzkaller/pkg/debugtracer"
	"github.com/google/syzkaller/sys/targets"
)

type testos struct {
	*git
}

var _ ConfigMinimizer = new(testos)

func newTestos(dir, tagGlob string, opts []RepoOpt) *testos {
	return &testos{
		git: newGit(dir, tagGlob, nil, opts),
	}
}

func (ctx *testos) PreviousReleaseTags(commit, compilerFamily string) ([]string, error) {
	return ctx.git.previousReleaseTags(commit, false, false, false)
}

func (ctx *testos) EnvForCommit(compilerFamily, binDir, commit string, kernelConfig []byte) (*BisectEnv, error) {
	return &BisectEnv{KernelConfig: kernelConfig}, nil
}

func (ctx *testos) Minimize(target *targets.Target, original, baseline []byte,
	dt debugtracer.DebugTracer, pred func(test []byte) (BisectResult, error)) ([]byte, error) {
	if res, err := pred(baseline); err != nil {
		return nil, err
	} else if res == BisectBad {
		return baseline, nil
	}
	switch string(baseline) {
	case "minimize-fails":
		return nil, fmt.Errorf("minimization failure")
	case "minimize-succeeds":
		config := []byte("new-minimized-config")
		pred(config)
		return config, nil
	default:
		return original, nil
	}
}
