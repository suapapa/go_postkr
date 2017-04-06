// Copyright 2017, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postkr

import (
	"encoding/xml"
	"errors"
	"fmt"
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

// Service represents epost.kr's post services
type Service struct {
	regkey       string
	lastQueryURL string

	totalCount, totalPage     int
	countPerPage, currentPage int
}

// NewService returns new Service. Call it with regkey.
func NewService(regkey string) *Service {
	s := Service{}
	s.regkey = regkey
	return &s
}
