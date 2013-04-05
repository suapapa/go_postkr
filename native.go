package postkr

import (
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"encoding/xml"
	"errors"
	"io"
)

func unmarshalCp949XML(data []byte, v interface{}) error {
	d := xml.NewDecoder(bytes.NewBuffer(data))
	d.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		if c != "cp949" && c != "euc-kr" {
			return nil, errors.New("unexpect charset: " + c)
		}
		tr, err := charset.TranslatorFrom("cp949")
		if err != nil {
			panic(err)
		}
		return charset.NewTranslatingReader(i, tr), nil
	}

	return d.Decode(v)
}

func encodeToCp949(utf8Str string) (string, error) {
	tr, err := charset.TranslatorTo("cp949")
	if err != nil {
		return "", err
	}
	_, r, err := tr.Translate([]byte(utf8Str), false)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
