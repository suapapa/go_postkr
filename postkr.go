// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	queryFmtStr = "http://biz.epost.go.kr/KpostPortal/openapi?" +
		"regkey=%s&target=%s&query=%s"
	fiveDigitQueryFmtStr = "http://biz.epost.go.kr/KpostPortal/openapi?" +
		"regkey=%s&target=%s&query=%s&countPerPage=%d&currentPage=%d"
)

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
	TotalCount int `xml:"pageinfo>totalCount"`
	TotalPage int `xml:"pageinfo>totalPage"`
	CountPerPage int `xml:"pageinfo>countPerPage"`
	CurrentPage int `xml:"pageinfo>currentPage"`
}

type Zipcode struct {
	XMLName xml.Name `xml:"item"`
	// Address: The address of the zipcode
	Address string `xml:"address"`
	// Code: The zip code number
	Code string `xml:"postcd"`
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
	totalCount int
	totalPage int
	countPerPage int
	currentPage int
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

func (s *Service) queryUrl(str string, target string) string {
	qs, err := encodeToCp949(str)
	if err != nil {
		// logE("iconv failed: ", err)
		return ""
	}

	query := url.QueryEscape(qs)
	s.lastQueryUrl = fmt.Sprintf(queryFmtStr, s.regkey, target, query)

	return s.lastQueryUrl
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
	if err := unmarshalCp949XML(body, &l); err != nil {
		var e serverError
		if err := unmarshalCp949XML(body, &e); err == nil {
			return nil, e.Error()
		}
		return nil, err
	}

	return l.Items, nil
}

func (s *Service) queryUrlOfFiveDigit(str string, target string, countPerPage int, currentPage int) string {
	qs, err := encodeToCp949(str)
	if err != nil {
		// logE("iconv failed: ", err)
		return ""
	}

	query := url.QueryEscape(qs)
	s.lastQueryUrl = fmt.Sprintf(fiveDigitQueryFmtStr, s.regkey, target, query, countPerPage, currentPage)

	return s.lastQueryUrl
}

// countPerPage : 페이지당 조회 건수
// currentPage : 조회할 페이지 번호
func (s *Service) SearchFiveDigitZipCode(key string, countPerPage int, currentPage int) ([]Zipcode, error) {
	if countPerPage < 10 {
		countPerPage = 10
	}
	if currentPage < 1 {
		currentPage = 1
	}
	
	url := s.queryUrlOfFiveDigit(key, "postNew", countPerPage, currentPage)
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
	if err := unmarshalCp949XML(body, &l); err != nil {
		var e serverError
		if err := unmarshalCp949XML(body, &e); err == nil {
			return nil, e.Error()
		}
		return nil, err
	}
	
	s.totalCount = l.TotalCount
	s.totalPage = l.TotalPage
	s.countPerPage = l.CountPerPage
	s.currentPage = l.CurrentPage

	return l.Items, nil
}

func (s *Service) TotalCount() int {
	return s.totalCount
}

func (s *Service) TotalPage() int {
	return s.totalPage
}

func (s *Service) CountPerPage() int {
	return s.countPerPage;
}

func (s *Service) CurrentPage() int {
	return s.currentPage
}
