// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Enumeration with an offset.
// Also includes a duplicate.

package main

import (
	"encoding"
	"fmt"
)

type Number int

const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One // Duplicate; note that AnotherOne doesn't appear below.
)

func main() {
	ck(One, "One")
	ck(Two, "Two")
	ck(Three, "Three")
	ck(AnotherOne, "One")
	ck(127, "Number(127)")
}

func ck(num Number, str string) {
	if fmt.Sprint(num) != str {
		panic("number.go: " + str)
	}

	var value Number
	u := interface{}(&value).(encoding.TextUnmarshaler)
	err := u.UnmarshalText([]byte(str))
	if num < One || num > Three {
		if err == nil || err.Error() != "Invalid Number: '"+str+"'" {
			panic("num.go: no error when it should")
		}
	} else {
		if err != nil {
			panic("num.go: " + err.Error())
		}
		if value != num {
			panic("num.go: " + str + " parsed as " + fmt.Sprint(num))
		}
	}
}
