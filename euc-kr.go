package postkr

import (
	"encoding/xml"
	"errors"
	"io"

	"github.com/suapapa/go_hangul/encoding/cp949"
)

func unmarshalCp949XML(r io.Reader, v interface{}) error {
	d := xml.NewDecoder(r)
	d.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		if c != "cp949" && c != "euc-kr" {
			return nil, errors.New("unexpect charset: " + c)
		}

		r, err := cp949.NewReader(i)
		if err != nil {
			return nil, err
		}

		return r, nil
	}

	return d.Decode(v)
}

func encodeToCp949(utf8Str string) (string, error) {
	cp949str, err := cp949.To([]byte(utf8Str))
	if err != nil {
		return "", err
	}

	return string(cp949str), nil
}
