// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*  Filename:    postkr.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-07-06 20:49:03.265118 +0900 KST
 *  Description: Main source file in postkr
 */

// Package postkr provides access to APIs of epost.kr, Korean post office.
//
//    http://biz.epost.go.kr/eportal/custom/custom_9.jsp?subGubun=sub_3&subGubun_1=cum_17&gubun=m07
//
// The epost.kr provide APIs for track snail-mail and search zip-code.
// But, I can't access to the tracking API with my key.
// So, It's not implemented currently.
//
// You need own key which is issued by eport.kr. Get one from this link:
//
//    http://biz.epost.go.kr/eportal/custom/custom_11.jsp?subGubun=sub_3&subGubun_1=cum_19&gubun=m07
//
package postkr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	iconv "github.com/djimenez/iconv-go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const queryFmtStr = "http://biz.epost.go.kr/KpostPortal/openapied?" +
	"regkey=%s&target=%s&query=%s"

type serverError struct {
	XMLName xml.Name `xml:"error"`
	Message string   `xml:"message"`
	Code    string   `xml:"error_code"`
}

func (e *serverError) String() string {
	return fmt.Sprintf("(%s) %s", e.Code, e.Message)
}

func (e *serverError) Error() error {
	return errors.New(e.String())
}

type zipcodeList struct {
	XMLName xml.Name  `xml:"post"`
	Items   []Zipcode `xml:"itemlist>item"`
}

type Zipcode struct {
	XMLName xml.Name `xml:"item"`
	// Address: The address of the zipcode
	Address string   `xml:"address"`
	// Code: The zip code number
	Code    string   `xml:"postcd"`
}

// String get string of Zipcode in form of "XXXXXX:Address of XXXXXX"
func (p *Zipcode) String() string {
	if p == nil {
		return "nil"
	}
	return p.Code + ":" + p.Address
}

// Codenum get zip code in uint
func (p *Zipcode) Codenum() uint {
	n, _ := strconv.ParseUint(p.Code, 10, 64)
	return uint(n)
}

type Service struct {
	regkey       string
	lastQueryUrl string
}

func (s *Service) queryUrl(str string, target string) string {
	iconvUtf8ToCp949, err := iconv.NewConverter("utf8", "cp949")
	if err != nil {
		log.Println("failed create iconv converter: ", err)
		return ""
	}
	defer iconvUtf8ToCp949.Close()

	qs, err := iconvUtf8ToCp949.ConvertString(str)
	if err != nil {
		log.Println("iconv failed: ", err)
		return ""
	}

	query := url.QueryEscape(qs)
	s.lastQueryUrl = fmt.Sprintf(queryFmtStr, s.regkey, target, query)

	return s.lastQueryUrl
}

func unmarshalCp949(data []byte, v interface{}) error {
	d := xml.NewDecoder(bytes.NewBuffer(data))
	d.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		if c != "euc-kr" && c != "cp949" {
			return nil, errors.New("unexpect charset: " + c)
		}
		return iconv.NewReader(i, "cp949", "utf8")
	}

	return d.Decode(v)
}

// SearchZipCode get Zipcodes for given [읍면동교] of an address.
func (s *Service) SerchZipCode(key string) ([]Zipcode, error) {
	url := s.queryUrl(key, "post")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", "ko")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var l zipcodeList
	if err := unmarshalCp949(body, &l); err != nil {
		var e serverError
		if err := unmarshalCp949(body, &e); err == nil {
			return nil, e.Error()
		}
		return nil, err
	}

	return l.Items, nil
}

// Initialize an new Service. Your own key of epost.kr open api is mandatory.
func NewService(regkey string) *Service {
	// if len(regkey) != 30 {
	// 	return nil
	// }
	s := new(Service)
	s.regkey = regkey
	return s
}
