// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	queryFmtStr = "http://biz.epost.go.kr/KpostPortal/openapi?" +
		"regkey=%s&target=%s&query=%s"
)

func (s *Service) queryURL(str string, target string) (string, error) {
	qs, err := encodeToCp949(str)
	if err != nil {
		return "", err
	}

	query := url.QueryEscape(qs)
	s.lastQueryURL = fmt.Sprintf(queryFmtStr, s.regkey, target, query)

	return s.lastQueryURL, nil
}

// SearchZipCode get Zipcodes for given [읍면동교] of an address.
func (s *Service) SearchZipCode(key string) ([]Zipcode, error) {
	url, err := s.queryURL(key, "post")
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

	return l.Items, nil
}
