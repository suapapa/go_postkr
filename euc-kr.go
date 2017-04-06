package postkr

import (
	"encoding/xml"
	"errors"
	"io"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

var (
	euckrEnc = korean.EUCKR.NewEncoder()
	euckrDec = korean.EUCKR.NewDecoder()
)

func unmarshalCp949XML(r io.Reader, v interface{}) error {
	d := xml.NewDecoder(r)
	d.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		if c != "cp949" && c != "euc-kr" {
			return nil, errors.New("unexpect charset: " + c)
		}

		r := transform.NewReader(i, euckrDec)
		return r, nil
	}

	return d.Decode(v)
}

// encodeToCp949 is used to transform input string to cp949 encoding
func encodeToCp949(utf8Str string) (string, error) {
	cp949str, _, err := transform.String(euckrEnc, utf8Str)
	if err != nil {
		return "", err
	}

	return cp949str, nil
}
