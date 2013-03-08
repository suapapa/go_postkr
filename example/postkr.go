// +build !appengine

// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/suapapa/go_postkr"
	"os"
)

var (
	POSTKR_APIKEY = os.Getenv("POSTKR_APIKEY")
)

func main() {
	s := postkr.NewService(POSTKR_APIKEY)
	l, _ := s.SerchZipCode(os.Args[1])

	for _, z := range l {
		fmt.Printf("%s - %s\n", z.Code, z.Address)
	}
}
