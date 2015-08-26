// +build !appengine

// Copyright 2013, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/jaehoonkim/go_postkr"
	"os"
)

var (
	POSTKR_APIKEY = os.Getenv("POSTKR_APIKEY")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s dong...", os.Args[0])
		os.Exit(1)
	}
	s := postkr.NewService(POSTKR_APIKEY)

	for _, d := range os.Args[1:] {
		fmt.Printf("Searching zipcode for %s...\n", d)
		l, _ := s.SearchFiveDigitZipCode(d, 10, 2)
		fmt.Println("TotalCount : ", s.TotalCount())
		fmt.Println("TotalPage : ", s.TotalPage())
		fmt.Println("CountPerPage : ", s.CountPerPage())
		fmt.Println("CurrentPage : ", s.CurrentPage())
		for _, z := range l {
			fmt.Printf("%s - %s\n", z.Code, z.Address)
		}
		fmt.Println("")
	}
}
