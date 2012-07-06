// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*  Filename:    postkr.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-07-06 20:49:03.265118 +0900 KST
 *  Description: Main source file in postkr
 */

// Package postkr does ....
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
)

type serverError struct {
	XMLName xml.Name `xml:"error"`
	Message string   `xml:"message"`
	Code    string   `xml:"error_code"`
}

type zipcodeList struct {
	XMLName xml.Name  `xml:"post"`
	Items   []Zipcode `xml:"itemlist>item"`
}

type Zipcode struct {
	XMLName xml.Name `xml:"item"`
	Address string   `xml:"address"`
	Code    string   `xml:"postcd"`
}

func (p *Zipcode) String() string {
	if p == nil {
		return "nil"
	}
	return p.Code + ":" + p.Address
}

var queryFmtStr string = "http://biz.epost.go.kr/KpostPortal/openapied?" +
	"regkey=%s&target=%s&query=%s"

type Service struct {
	regkey       string // 사용신청을 통해 받은 인증 key 스트링(30자리)
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
			errStr := fmt.Sprintf("(%s) %s", e.Code, e.Message)
			return nil, errors.New(errStr)
		}
		return nil, err
	}

	return l.Items, nil
}

func NewService(regkey string) *Service {
	s := new(Service)
	s.regkey = regkey
	return s
}
