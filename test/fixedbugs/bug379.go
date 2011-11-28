// errchk $G $D/$F.go

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Issue 2452.

// Check that the error messages says 
//	bug378.go:17: 3 not used
// and not
//	bug378.go:17: 1 not used

package main

func main() {
	1 + 2 // ERROR "3 not used|value computed is not used"
}
