// errchk $G -e $D/$F.go

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Tbyte []byte
type Trune []rune
type Tint64 []int64
type Tstring string

func main() {
	s := "hello"
	sb := []byte("hello")
	sr := []rune("hello")
	si := []int64{'h', 'e', 'l', 'l', 'o'}

	ts := Tstring(s)
	tsb := Tbyte(sb)
	tsr := Trune(sr)
	tsi := Tint64(si)

	_ = string(s)
	_ = []byte(s)
	_ = []rune(s)
	_ = []int64(s) // ERROR "cannot convert.*\[\]int64"
	_ = Tstring(s)
	_ = Tbyte(s)
	_ = Trune(s)
	_ = Tint64(s) // ERROR "cannot convert.*Tint64"

	_ = string(sb)
	_ = []byte(sb)
	_ = []rune(sb)  // ERROR "cannot convert.*\[\]rune"
	_ = []int64(sb) // ERROR "cannot convert.*\[\]int64"
	_ = Tstring(sb)
	_ = Tbyte(sb)
	_ = Trune(sb)  // ERROR "cannot convert.*Trune"
	_ = Tint64(sb) // ERROR "cannot convert.*Tint64"

	_ = string(sr)
	_ = []byte(sr) // ERROR "cannot convert.*\[\]byte"
	_ = []rune(sr)
	_ = []int64(sr) // ERROR "cannot convert.*\[\]int64"
	_ = Tstring(sr)
	_ = Tbyte(sr) // ERROR "cannot convert.*Tbyte"
	_ = Trune(sr)
	_ = Tint64(sr) // ERROR "cannot convert.*Tint64"

	_ = string(si) // ERROR "cannot convert.* string"
	_ = []byte(si) // ERROR "cannot convert.*\[\]byte"
	_ = []rune(si) // ERROR "cannot convert.*\[\]rune"
	_ = []int64(si)
	_ = Tstring(si) // ERROR "cannot convert.*Tstring"
	_ = Tbyte(si)   // ERROR "cannot convert.*Tbyte"
	_ = Trune(si)   // ERROR "cannot convert.*Trune"
	_ = Tint64(si)

	_ = string(ts)
	_ = []byte(ts)
	_ = []rune(ts)
	_ = []int64(ts) // ERROR "cannot convert.*\[\]int64"
	_ = Tstring(ts)
	_ = Tbyte(ts)
	_ = Trune(ts)
	_ = Tint64(ts) // ERROR "cannot convert.*Tint64"

	_ = string(tsb)
	_ = []byte(tsb)
	_ = []rune(tsb)  // ERROR "cannot convert.*\[\]rune"
	_ = []int64(tsb) // ERROR "cannot convert.*\[\]int64"
	_ = Tstring(tsb)
	_ = Tbyte(tsb)
	_ = Trune(tsb)  // ERROR "cannot convert.*Trune"
	_ = Tint64(tsb) // ERROR "cannot convert.*Tint64"

	_ = string(tsr)
	_ = []byte(tsr) // ERROR "cannot convert.*\[\]byte"
	_ = []rune(tsr)
	_ = []int64(tsr) // ERROR "cannot convert.*\[\]int64"
	_ = Tstring(tsr)
	_ = Tbyte(tsr) // ERROR "cannot convert.*Tbyte"
	_ = Trune(tsr)
	_ = Tint64(tsr) // ERROR "cannot convert.*Tint64"

	_ = string(tsi) // ERROR "cannot convert.* string"
	_ = []byte(tsi) // ERROR "cannot convert.*\[\]byte"
	_ = []rune(tsi) // ERROR "cannot convert.*\[\]rune"
	_ = []int64(tsi)
	_ = Tstring(tsi) // ERROR "cannot convert.*Tstring"
	_ = Tbyte(tsi)   // ERROR "cannot convert.*Tbyte"
	_ = Trune(tsi)   // ERROR "cannot convert.*Trune"
	_ = Tint64(tsi)
}