// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"fmt"
	"os"
	/* "testing" */)

var (
	PostKrAPIKey = os.Getenv("POSTKR_APIKEY")
)

// func ExampleSearchZipCode() {
// 	s := NewService(PostKrAPIKey)
// 	l, _ := s.SearchZipCode("도곡동")
// 	z := l[0]
// 	fmt.Printf("%s - %s\n", z.Code, z.Address)
// 	// Output: 135270 - 서울 강남구 도곡동
// }

// func ExampleNumberCode() {
// 	s := NewService(PostKrAPIKey)
// 	l, _ := s.SearchZipCode("내곡동")
// 	cn := l[0].Codenum()
// 	fmt.Printf("%03d-%03d\n", cn/1000, cn%1000)
// 	// Output: 137-180
// }

// func ExampleSearchZipCodeForUnexistDong() {
// 	s := NewService(PostKrAPIKey)
// 	l, _ := s.SearchZipCode("어우동")
// 	fmt.Println(len(l))
// 	// Output: 0
// }

func ExampleSearchZipCode5() {
	s := NewService(PostKrAPIKey)
	l, _ := s.SearchZipCode5("선릉로112길 49", 10, 1)
	z := l[0]
	fmt.Printf("%s - %s\n", z.Code, z.Address)
	// Output: 06097 - 서울특별시 강남구 선릉로112길 49 (삼성동)
}

func ExampleInvalidRegKey() {
	s := NewService("12345")
	_, e := s.SearchZipCode("dummy")
	fmt.Println(e)
	// Output: expected element type <post> but have <error>
}
