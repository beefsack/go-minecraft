package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"github.com/beefsack/go-minecraft/nbt"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Please specify a command")
	}
	command := flag.Args()[0]
	commandParams := flag.Args()[1:]
	switch command {
	case "nbt":
		f, err := os.Open(commandParams[0])
		if err != nil {
			fmt.Errorf("Error opening file: %s\n", err.Error())
			return
		}
		var br *bufio.Reader
		gzr, err := gzip.NewReader(f)
		if err == nil {
			// Compressed
			br = bufio.NewReader(gzr)
		} else {
			// Not compressed
			br = bufio.NewReader(f)
		}
		data, err := nbt.Decode(br)
		if err != nil {
			fmt.Errorf("Error decoding file: %s\n", err.Error())
			return
		}
		fmt.Printf("%#v\n", data)
	}
}
