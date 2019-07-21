package main

/* This is Rexa's registry: the file where you should register all thread types that would partake
in your software's life cycle. */

import (
	"github.com/qamarian-mop/rx-lib"

	// Import below, all thread types you wish to be part of the system
)

var threadTypes map[string]ThreadTypeRegister = map[string]ThreadTypeRegister { /* One more step to
	make a thread type a part of this system: register the thread type here. Key should be an ID
	for the thread type (value must be numeric string). Value of key should be the register of
	the thread type. */
}
