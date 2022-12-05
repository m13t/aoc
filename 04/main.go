package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Range struct {
	From int
	To   int
}

func (r1 Range) encloses(r2 Range) bool {
	if r1.From <= r2.From && r1.To >= r2.To {
		return true
	}

	return false
}

func (r1 Range) overlaps(r2 Range) bool {
	if (r1.From >= r2.From && r1.From <= r2.To) || (r1.To >= r2.From && r1.To <= r2.To) {
		return true
	}

	return false
}

type RangeSet [2]Range

func (r RangeSet) encloses() bool {
	return r[0].encloses(r[1]) || r[1].encloses(r[0])
}

func (r RangeSet) overlaps() bool {
	return r[0].overlaps(r[1]) || r[1].overlaps(r[0])
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := parseData(f)
	if err != nil {
		panic(err)
	}

	res, err := roundOne(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round One Result: %d\n", res)

	res, err = roundTwo(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round Two Result: %d\n", res)
}

func roundOne(data []RangeSet) (int, error) {
	var total int

	for _, set := range data {
		if set.encloses() {
			total += 1
		}
	}

	return total, nil
}

func roundTwo(data []RangeSet) (int, error) {
	var total int

	for _, set := range data {
		if set.overlaps() {
			total += 1
		}
	}

	return total, nil
}

func parseData(r io.Reader) ([]RangeSet, error) {
	s := bufio.NewScanner(r)

	reg := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)

	var ranges []RangeSet

	for s.Scan() {
		res := reg.FindStringSubmatch(s.Text())
		if len(res) != 5 {
			return nil, errors.New("unexpected number of digits in range entry")
		}

		s1, err := strconv.Atoi(res[1])
		if err != nil {
			return nil, err
		}
		e1, err := strconv.Atoi(res[2])
		if err != nil {
			return nil, err
		}
		r1 := Range{s1, e1}

		s2, err := strconv.Atoi(res[3])
		if err != nil {
			return nil, err
		}
		e2, err := strconv.Atoi(res[4])
		if err != nil {
			return nil, err
		}
		r2 := Range{s2, e2}

		ranges = append(ranges, RangeSet{r1, r2})
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return ranges, nil
}
