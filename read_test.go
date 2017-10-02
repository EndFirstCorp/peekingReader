package peekingReader

import (
	"errors"
)

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
