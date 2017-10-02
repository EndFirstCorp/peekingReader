package peekingReader

import (
	"io"
	"strings"
	"testing"
)

func TestBufReadBytes(t *testing.T) {
	r := NewBufReader(strings.NewReader("hello there my friend. How are you?"))
	b, _ := r.ReadBytes(31)
	if string(b) != "hello there my friend. How are " {
		t.Error("expected correct bytes", string(b))
	}
	if b, err := r.ReadBytes(4); err != nil || len(b) != 4 || string(b) != "you?" {
		t.Error("expected shorter string", err, string(b))
	}
	if _, err := r.ReadByte(); err != io.EOF {
		t.Fatal("expected error", err)
	}
}
