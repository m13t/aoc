package main

import (
	"bufio"
	"errors"
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

	bags, err := parseData(f)
	if err != nil {
		panic(err)
	}

	r1, err := roundOne(bags)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round One Total: %d\n", r1)

	r2, err := roundTwo(bags)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round Two Total: %d\n", r2)
}

func roundOne(bags [][]byte) (int, error) {
	var total int

	for _, bag := range bags {
		l := len(bag) / 2
		p1, p2 := bag[:l], bag[l:]
		i := intersect(p1, p2)

		if len(i) != 1 {
			return total, errors.New("unexpected item in the bagging area")
		}

		total += int(i[0])
	}

	return total, nil
}

func roundTwo(bags [][]byte) (int, error) {
	var total int

	for x := 0; x < len(bags); x += 3 {
		i := intersect(intersect(bags[x], bags[x+1]), bags[x+2])
		if len(i) != 1 {
			return total, errors.New("unexpected item in the bagging area")
		}

		total += int(i[0])
	}

	return total, nil
}

func contains(a []byte, b byte) bool {
	for _, x := range a {
		if x == b {
			return true
		}
	}

	return false
}

func intersect(a, b []byte) []byte {
	var out []byte

	for _, x := range a {
		if contains(b, x) && !contains(out, x) {
			out = append(out, x)
		}
	}

	return out
}

func parseData(r io.Reader) ([][]byte, error) {
	var out [][]byte

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		tb := s.Bytes()
		buf := make([]byte, len(tb))
		copy(buf, tb)

		for x, b := range buf {
			if b >= 97 && b <= 122 {
				buf[x] -= 96
				continue
			}

			if b >= 65 && b <= 90 {
				buf[x] -= 38
				continue
			}

			return nil, errors.New("invalid character found in sequence")
		}

		out = append(out, buf)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
