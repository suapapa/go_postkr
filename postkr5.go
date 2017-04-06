// Copyright 2017, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	queryFmtStr5 = "http://biz.epost.go.kr/KpostPortal/openapi?" +
		"regkey=%s&target=%s&query=%s&countPerPage=%d&currentPage=%d"
)

func (s *Service) queryURL5(str string, target string,
	countPerPage, currentPage int,
) (string, error) {
	qs, err := encodeToCp949(str)
	if err != nil {
		return "", err
	}

	query := url.QueryEscape(qs)
	s.lastQueryURL = fmt.Sprintf(queryFmtStr5, s.regkey, target, query, countPerPage, currentPage)

	return s.lastQueryURL, nil
}

// SearchZipCode5 search new format of zipcode
// with road based address
//  * countPerPage : 페이지당 조회 건수
//  * currentPage : 조회할 페이지 번호
func (s *Service) SearchZipCode5(key string, countPerPage int, currentPage int) ([]Zipcode, error) {
	if countPerPage < 10 {
		countPerPage = 10
	}
	if currentPage < 1 {
		currentPage = 1
	}

	url, err := s.queryURL5(key, "postNew", countPerPage, currentPage)
	if err != nil {
		return nil, err
	}

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

	// body, err := ioutil.ReadAll(resp.Body)
	var l zipcodeList
	if err := unmarshalCp949XML(resp.Body, &l); err != nil {
		var e serverError
		if err := unmarshalCp949XML(resp.Body, &e); err == nil {
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
	return s.countPerPage
}

func (s *Service) CurrentPage() int {
	return s.currentPage
}
