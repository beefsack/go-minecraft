package nbt

import (
	"fmt"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("../test_files/world/players/beefsack.dat")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	data, err := Decode(f)
	fmt.Println("%#v\n", data)
	if err != nil {
		t.Fatal(err)
	}
}
