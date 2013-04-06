//
// Written by Maxim Khitrov (February 2013)
//

package sqlite3

/*
#include "sqlite3.h"

#cgo CFLAGS: -DNDEBUG -fno-stack-check -fno-stack-protector -mno-stack-arg-probe

int shell_main(int, void*);
*/
import "C"

import "unsafe"

// shell runs shell_main defined in shell_windows.c (renamed main).
func shell(args []string) int {
	// Copy all arguments into a single []byte, terminating each one with '\0'
	buf := make([]byte, 0, 256)
	for _, arg := range args {
		buf = append(append(buf, arg...), 0)
	}

	// Fill argv with pointers to the start of each null-terminated string
	argv := make([]uintptr, len(args))
	base := uintptr(cBytes(buf))
	for i, arg := range args {
		argv[i] = base
		base += uintptr(len(arg)) + 1
	}
	return int(C.shell_main(C.int(len(args)), unsafe.Pointer(&argv[0])))
}

// errstr uses the native implementation of sqlite3_errstr.
func errstr(rc C.int) string {
	return C.GoString(C.sqlite3_errstr(rc))
}