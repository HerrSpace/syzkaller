// Copyright 2021 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

#define AK_DONT_REPLACE_STD 1

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>
#include <unistd.h>
#include <syscall.h>

#include "nocover.h"

static void os_init(int argc, char** argv, void* data, size_t data_size)
{
	int prot = PROT_READ | PROT_WRITE;
	int flags = MAP_ANON | MAP_PRIVATE | MAP_FIXED;
	if (mmap(data, data_size, prot, flags, -1, 0) != data)
		fail("mmap of data segment failed");
}

static intptr_t execute_syscall(const call_t* c, intptr_t a[kMaxArgs])
{
	return syscall4(c->sys_nr, a[0], a[1], a[2], a[3]);
}
