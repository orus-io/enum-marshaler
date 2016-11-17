// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Import "C" shouldn't be imported.

package main

/*
#define HELLO 1
*/
import "C"

import (
	"encoding"
	"fmt"
)

type Cgo uint32

const (
	// MustScanSubDirs indicates that events were coalesced hierarchically.
	MustScanSubDirs Cgo = 1 << iota
)

func main() {
	_ = C.HELLO
	ck(MustScanSubDirs, "MustScanSubDirs")
}

func ck(cgo Cgo, str string) {
	if fmt.Sprint(cgo) != str {
		panic("cgo.go: " + str)
	}

	var value Cgo
	u := interface{}(&value).(encoding.TextUnmarshaler)
	err := u.UnmarshalText([]byte(str))
	if cgo != MustScanSubDirs {
		if err == nil || err.Error() != "Invalid Cgo: '"+str+"'" {
			panic("day.go: no error when it should")
		}
	} else {
		if err != nil {
			panic("cgo.go: " + err.Error())
		}
		if value != cgo {
			panic("cgo.go: " + str + " parsed as " + fmt.Sprint(cgo))
		}
	}
}
