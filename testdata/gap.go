// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gaps and an offset.

package main

import (
	"encoding"
	"fmt"
)

type Gap int

const (
	Two    Gap = 2
	Three  Gap = 3
	Five   Gap = 5
	Six    Gap = 6
	Seven  Gap = 7
	Eight  Gap = 8
	Nine   Gap = 9
	Eleven Gap = 11
)

func main() {
	ck(0, "Gap(0)")
	ck(1, "Gap(1)")
	ck(Two, "Two")
	ck(Three, "Three")
	ck(4, "Gap(4)")
	ck(Five, "Five")
	ck(Six, "Six")
	ck(Seven, "Seven")
	ck(Eight, "Eight")
	ck(Nine, "Nine")
	ck(10, "Gap(10)")
	ck(Eleven, "Eleven")
	ck(12, "Gap(12)")
}

func ck(gap Gap, str string) {
	if fmt.Sprint(gap) != str {
		panic("gap.go: " + str)
	}

	var value Gap
	u := interface{}(&value).(encoding.TextUnmarshaler)
	err := u.UnmarshalText([]byte(str))
	if str[:3] == "Gap" {
		if err == nil || err.Error() != "Invalid Gap: '"+str+"'" {
			panic("gap.go: no error when it should")
		}
	} else {
		if err != nil {
			panic("gap.go: " + err.Error())
		}
		if value != gap {
			panic("gap.go: " + str + " parsed as " + fmt.Sprint(gap))
		}
	}
}
