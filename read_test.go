package peekingReader

import (
	"errors"
	"io"
	"testing"
)

func TestSkipSubsequent(t *testing.T) {
	r := NewMemReader([]byte("  \t   there"))
	SkipSubsequent(r, []byte{' ', '\t'})
	if b, _ := r.ReadByte(); b != 't' {
		t.Error("Expected correct value", b)
	}

	r = NewMemReader([]byte("     "))
	_, err := SkipSubsequent(r, []byte{' ', '\t'})
	if err != io.EOF {
		t.Error("expected EOF error")
	}
}

func TestSkipSpaces(t *testing.T) {
	r := NewMemReader([]byte("  \t   there"))
	SkipSpaces(r)
	if b, _ := r.ReadByte(); b != 't' {
		t.Error("Expected correct value", b)
	}
}

func TestReadUntil(t *testing.T) {
	r := NewMemReader([]byte("  \t   there"))
	if b, _ := ReadUntil(r, 'e'); string(b) != "  \t   th" {
		t.Error("expected valid value", string(b))
	}
}

func TestReadUntilAny(t *testing.T) {
	r := NewMemReader([]byte("  \t   there"))
	if b, _ := ReadUntilAny(r, []byte{'t', 'e'}); string(b) != "  \t   " {
		t.Error("expected valid value", string(b))
	}

	r = NewMemReader([]byte("  \t   there"))
	if _, err := ReadUntilAny(r, []byte{'a', 'b'}); err != io.EOF {
		t.Error("expected error", err)
	}
}

type erroringPeeker struct{}

func (p *erroringPeeker) Peek(n int) ([]byte, error) {
	return nil, errors.New("fail")
}

func (p erroringPeeker) ReadByte() (byte, error) {
	return '\x00', nil
}

func (p erroringPeeker) ReadBytes(size int) ([]byte, error) {
	return nil, nil
}
