package nbt

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	data, err := Decode(bigTestReader())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Decoded data.")
	buf := bytes.NewBuffer([]byte{})
	err = Encode(data, buf)
	if err != nil {
		t.Fatal(err)
	}
	decodedData, err := Decode(buf)
	if err != nil {
		t.Fatal(err)
	}
	testMatchesBigTest(t, decodedData)
}

func TestDetectType(t *testing.T) {
	tagType, err := DetectPayloadType(byte(5))
	if err != nil {
		t.Fatal(err)
	}
	if tagType != TAG_BYTE {
		t.Fatal("Did not detect byte, got", tagType)
	}
}
