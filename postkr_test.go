// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

/*  Filename:    postkr_test.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-07-06 20:49:03.265612 +0900 KST
 *  Description: Main test file for postkr
 */

import (
	"fmt"
	/* "testing" */
)

const postkrRegKey = "5e12d7ed7799470b81298981375429"

func ExampleSearchZipCode() {
	s := NewService(postkrRegKey)
	l, _ := s.SerchZipCode("도곡동")
	z := l[0]
	fmt.Printf("%s - %s\n", z.Code, z.Address)
	// Output: 135270 - 서울 강남구 도곡동
}

func ExampleSearchZipCodeForUnexistDong() {
	s := NewService(postkrRegKey)
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
