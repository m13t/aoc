package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/m13t/aoc/05/types"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	inventory, instructions, err := parseData(f)
	if err != nil {
		panic(err)
	}

	out, err := roundOne(inventory, instructions)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round One Output: %s\n", out)

	out, err = roundTwo(inventory, instructions)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round Two Output: %s\n", out)
}

func roundOne(inv [][]byte, ins [][]int) (string, error) {
	stacks := types.NewStacks()
	for x := len(inv) - 1; x >= 0; x-- {
		for ix, item := range inv[x] {
			if item != 0 {
				stacks.PushTo(ix, item)
			}
		}
	}

	for _, i := range ins {
		stacks.MoveOne(i[0], i[1], i[2])
	}

	res := make([]byte, len(stacks))
	for x, stack := range stacks {
		if v, ok := stack.Peek(); ok {
			res[x] = v
		}
	}

	return string(res), nil
}

func roundTwo(inv [][]byte, ins [][]int) (string, error) {
	stacks := types.NewStacks()
	for x := len(inv) - 1; x >= 0; x-- {
		for ix, item := range inv[x] {
			if item != 0 {
				stacks.PushTo(ix, item)
			}
		}
	}

	for _, i := range ins {
		stacks.MoveSet(i[0], i[1], i[2])
	}

	res := make([]byte, len(stacks))
	for x, stack := range stacks {
		if v, ok := stack.Peek(); ok {
			res[x] = v
		}
	}

	return string(res), nil
}

func parseData(r io.Reader) ([][]byte, [][]int, error) {
	s := bufio.NewScanner(r)

	// Start by reading header
	headerMode := true

	// Crate output variables
	var rows [][]byte
	var instructions [][]int

	// Process each line
	for s.Scan() {
		line := s.Bytes()

		if headerMode {
			if len(line) == 0 {
				headerMode = false
				continue
			}

			row, err := parseRow(line)
			if err != nil {
				return nil, nil, err
			}
			if row != nil {
				rows = append(rows, row)
			}
		} else {
			instruction, err := parseInstruction(line)
			if err != nil {
				return nil, nil, err
			}
			if instruction != nil {
				instructions = append(instructions, instruction)
			}
		}
	}

	if err := s.Err(); err != nil {
		return nil, nil, err
	}

	return rows, instructions, nil
}

func parseRow(b []byte) ([]byte, error) {
	var out []byte

	for x := 0; x < len(b)+1; x += 4 {
		if b[x+1] == ' ' {
			out = append(out, 0)
			continue
		}

		if b[x+1] >= 'A' && b[x+1] <= 'Z' {
			out = append(out, b[x+1])
			continue
		}
	}

	return out, nil
}

func parseInstruction(b []byte) ([]int, error) {
	parts := strings.Split(string(b), " ")

	p1, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	p2, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}

	p3, err := strconv.Atoi(parts[5])
	if err != nil {
		return nil, err
	}

	return []int{p1, p2, p3}, nil
}
