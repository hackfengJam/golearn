package cast

import (
	"bytes"
	"testing"
)

func BenchmarkStringToBytes(b *testing.B) {
	s := "hello world"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		StringToBytes(s)
	}
}

func TestLen(t *testing.T) {
	n := len("我是😄🤑")
	if n != 14 {
		t.Error(n)
	}
	n = len([]byte("我是😁😝"))
	if n != 14 {
		t.Error(n)
	}
}

func TestStringToBytes(t *testing.T) {
	s := string([]byte("hello"))
	b := StringToBytes(s)
	if !bytes.Equal(b, []byte(s)) {
		t.Errorf("%s", b)
	}
	b[0] = 'z'
	if s != "zello" {
		t.Error(s)
	}
}

func BenchmarkBytesToString(b *testing.B) {
	s := []byte("hello world")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		BytesToString(s)
	}
}

func TestBytesToString(t *testing.T) {
	b := []byte("hello")
	s := BytesToString(b)
	if s != string(b) {
		t.Error(s)
	}
	b[1] = 'z'
	if s != "hzllo" {
		t.Error(s)
	}
}
