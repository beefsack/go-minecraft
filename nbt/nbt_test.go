package nbt

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("../test_files/world/players/beefsack.dat")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	_, err = Decode(f)
	if err != nil {
		t.Fatal(err)
	}
}
