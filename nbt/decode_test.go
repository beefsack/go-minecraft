package nbt

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("../test_files/bigtest.nbt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	data, err := Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Decoded data.")
	testMatchesBigTest(t, data)
}
