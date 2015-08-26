// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"fmt"
	"os"
	/* "testing" */)

var (
	POSTKR_APIKEY = os.Getenv("POSTKR_APIKEY")
)

func ExampleSearchZipCode() {
	s := NewService(POSTKR_APIKEY)
	l, _ := s.SerchZipCode("도곡동")
	z := l[0]
	fmt.Printf("%s - %s\n", z.Code, z.Address)
	// Output: 135270 - 서울 강남구 도곡동
}

func ExampleNumberCode() {
	s := NewService(POSTKR_APIKEY)
	l, _ := s.SerchZipCode("내곡동")
	cn := l[0].Codenum()
	fmt.Printf("%03d-%03d\n", cn/1000, cn%1000)
	// Output: 137-180
}

func ExampleSearchZipCodeForUnexistDong() {
	s := NewService(POSTKR_APIKEY)
	l, _ := s.SerchZipCode("어우동")
	fmt.Println(len(l))
	// Output: 0
}

func ExampleInvalidRegKey() {
	s := NewService("12345")
	_, e := s.SerchZipCode("dummy")
	fmt.Println(e)
	// Output: (ERR-002) 우편번호 조회의 등록되지 않거나 유효하지 않은 인증key 입니다.
}
