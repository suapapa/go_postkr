// Copyright 2017, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"encoding/xml"
	"strconv"
)

type zipcodeList struct {
	XMLName      xml.Name  `xml:"post"`
	Items        []Zipcode `xml:"itemlist>item"`
	TotalCount   int       `xml:"pageinfo>totalCount"`
	TotalPage    int       `xml:"pageinfo>totalPage"`
	CountPerPage int       `xml:"pageinfo>countPerPage"`
	CurrentPage  int       `xml:"pageinfo>currentPage"`
}

// Zipcode represents snailmail's zip code
type Zipcode struct {
	XMLName xml.Name `xml:"item"`
	// Address: The address of the zipcode
	Address string `xml:"address"`
	// Code: The zip code number
	Code string `xml:"postcd"`
}

// String get string of Zipcode in form of "XXXXXX:Address of XXXXXX"
func (p Zipcode) String() string {
	return p.Code + ":" + p.Address
}

// Codenum get zip code in uint
func (p Zipcode) Codenum() uint {
	n, _ := strconv.ParseUint(p.Code, 10, 64)
	return uint(n)
}
