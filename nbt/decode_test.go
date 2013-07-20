package nbt

import (
	"bufio"
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	data, err := Decode(bigTestReader())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Decoded data.")
	testMatchesBigTest(t, data)
}

func TestDecodeIntArray(t *testing.T) {
	// 3, 1, 2, 3
	buf := []byte{0, 0, 0, 3, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3}
	r := bytes.NewReader(buf)
	br := bufio.NewReader(r)
	intArray, err := DecodeIntArray(br)
	if err != nil {
		t.Fatal(err)
	}
	if len(intArray) != 3 {
		t.Fatal("Array wrong length, got", len(intArray))
	}
	if intArray[0] != 1 {
		t.Fatal("Element 0 incorrect, got", intArray[0])
	}
	if intArray[1] != 2 {
		t.Fatal("Element 1 incorrect, got", intArray[1])
	}
	if intArray[2] != 3 {
		t.Fatal("Element 2 incorrect, got", intArray[2])
	}
}
