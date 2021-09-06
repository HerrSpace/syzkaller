// Copyright 2021 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package report

import (
	"regexp"
)

func ctorSerenity(cfg *config) (reporterImpl, []string, error) {
	symbolizeRes := []*regexp.Regexp{}
	ctx, err := ctorBSD(cfg, serenityOopses, symbolizeRes)
	return ctx, nil, err
}

var serenityOopses = append([]*oops{
	{
		[]byte("ASSERTION FAILED:"),
		[]oopsFormat{
			{
				title: compile("ASSERTION FAILED: (.+)"),
				fmt:   "panic: %[1]v",
			},
		},
		[]*regexp.Regexp{},
	},
}, commonOopses...)
