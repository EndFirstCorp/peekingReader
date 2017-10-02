package peekingReader

import (
	"io"
	"testing"
)

func TestMemReadBytes(t *testing.T) {
	r := NewMemReader([]byte("hello"))
	b, _ := r.ReadBytes(3)
	if string(b) != "hel" {
		t.Error("expected correct bytes")
	}
	if _, err := r.ReadBytes(3); err != io.EOF {
		t.Error("expected EOF", err)
	}
}
