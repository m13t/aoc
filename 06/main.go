package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := parseData(f); err != nil {
		panic(err)
	}
}

func parseData(r io.Reader) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		buf := s.Bytes()

		for x := 4; x < len(buf); x++ {
			w := buf[x-4 : x]

			if !containsDuplicates(w) {
				fmt.Printf("Packet Marker %q at %d\n", w, x)
				break
			}
		}

		for x := 14; x < len(buf); x++ {
			w := buf[x-14 : x]

			if !containsDuplicates(w) {
				fmt.Printf("Message Marker %q at %d\n", w, x)
				break
			}
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}

func containsDuplicates(buf []byte) bool {
	out := map[byte]struct{}{}

	for _, v := range buf {
		if _, ok := out[v]; !ok {
			out[v] = struct{}{}
		}
	}

	return len(out) < len(buf)
}
