package conversion

import (
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
	t.Logf("transform before : %s , after : %s", buff, res)
}
