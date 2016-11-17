// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Signed integers spanning zero.

package main

import (
	"encoding"
	"fmt"
)

type Num int

const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)

func main() {
	ck(-3, "Num(-3)")
	ck(m_2, "m_2")
	ck(m_1, "m_1")
	ck(m0, "m0")
	ck(m1, "m1")
	ck(m2, "m2")
	ck(3, "Num(3)")
}

func ck(num Num, str string) {
	if fmt.Sprint(num) != str {
		panic("num.go: " + str)
	}

	var value Num
	u := interface{}(&value).(encoding.TextUnmarshaler)
	err := u.UnmarshalText([]byte(str))
	if num < m_2 || num > m2 {
		if err == nil || err.Error() != "Invalid Num: '"+str+"'" {
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
