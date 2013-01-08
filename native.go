package postkr

import (
	"bytes"
	"encoding/xml"
	"errors"
	iconv "github.com/djimenez/iconv-go"
	"io"
)

func unmarshalCp949XML(data []byte, v interface{}) error {
	d := xml.NewDecoder(bytes.NewBuffer(data))
	d.CharsetReader = func(c string, i io.Reader) (io.Reader, error) {
		if c != "euc-kr" && c != "cp949" {
			return nil, errors.New("unexpect charset: " + c)
		}
		return iconv.NewReader(i, "cp949", "utf8")
	}

	return d.Decode(v)
}

func encodeToCp949(utf8Str string) (string, error) {
	return iconv.ConvertString(utf8Str, "utf8", "cp949")
}
