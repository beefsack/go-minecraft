package nbt

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
)

const (
	MODE_OUTSIDE_TAG = iota
)

func Decode(r io.Reader) (map[string]interface{}, error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()
	br := bufio.NewReader(gzr)
	for {
		b, err := br.ReadByte()
		fmt.Print(string([]byte{b}))
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return nil, nil
}
