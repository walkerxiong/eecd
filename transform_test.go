package cced

import (
	"github.com/gbabyX/cced/charset/utf16"
	"io/ioutil"
	"testing"
)

func TestGbkToUtf8(t *testing.T) {
	buff, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	res, err := GbkToUtf8(buff)
	if err != nil {
		t.Fatal(err)
	}
	// e4b8a5e88283e79a84e8afb4efbc8ce68891e698afe4b880e4b8aa4368696e657365efbc8ce68891e69c89e4b8ade59bbde9ad82e3808263616e207520756e6465727374616e643f3f3f
	t.Logf("%x \n", res)
	r, err := Utf8ToUtf16(res)
	//4e25808376848bf4ff0c6211662f4e004e2a004300680069006e006500730065ff0c621167094e2d56fd9b42300200630061006e0020007500200075006e006400650072007300740061006e0064003f003f003f
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%x \n", r)
	// 4e25808376848bf4ff0c6211662f4e004e2a4368696e657365ff0c621167094e2d56fd9b42300263616e207520756e6465727374616e643f3f3f
	r2, err := Utf16ToUtf8(r)
	t.Logf("%x \n", r2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s \n", r2)
}

func TestUtf8ToUtf16(t *testing.T) {
	buff, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x \n", buff)
	r1, err := utf16.Decode(buff)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x \n", r1)
}
