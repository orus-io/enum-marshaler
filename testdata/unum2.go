// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Unsigned integers - check maximum size

package main

import (
	"encoding"
	"fmt"
)

type Unum2 uint8

const (
	Zero Unum2 = iota
	One
	Two
)

func main() {
	ck(Zero, "Zero")
	ck(One, "One")
	ck(Two, "Two")
	ck(3, "Unum2(3)")
	ck(255, "Unum2(255)")
}

func ck(unum Unum2, str string) {
	if fmt.Sprint(unum) != str {
		panic("unum.go: " + str)
	}

	var value Unum2
	u := interface{}(&value).(encoding.TextUnmarshaler)
	err := u.UnmarshalText([]byte(str))
	if unum > Two {
		if err == nil || err.Error() != "Invalid Unum2: '"+str+"'" {
			panic("unum.go: no error when it should")
		}
	} else {
		if err != nil {
			panic("unum.go: " + err.Error())
		}
		if value != unum {
			panic("unum.go: " + str + " parsed as " + fmt.Sprint(unum))
		}
	}
}
