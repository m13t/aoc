package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	items []int
}

func newElf() *elf {
	return &elf{
		items: []int{},
	}
}

func (e *elf) AddItem(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	e.items = append(e.items, i)

	return nil
}

func (e *elf) Total() int {
	var total int
	for _, x := range e.items {
		total += x
	}
	return total
}

type Elves []elf

func (e Elves) Len() int {
	return len(e)
}

func (e Elves) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Elves) Less(i, j int) bool {
	return e[i].Total() < e[j].Total()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a line scanner on the file
	s := bufio.NewScanner(f)

	// All our elves
	var elves Elves

	// Current elf pointer
	var currentElf *elf

	for s.Scan() {
		// Lazily set the first elf
		if currentElf == nil {
			currentElf = newElf()
		}

		t := s.Text()
		if t == "" {
			elves = append(elves, *currentElf)
			currentElf = newElf()
		} else {
			if err := currentElf.AddItem(t); err != nil {
				panic(err)
			}
		}
	}

	// Check if there was an error scanning the input
	if err := s.Err(); err != nil {
		panic(err)
	} else {
		// If no error, append the last elf
		elves = append(elves, *currentElf)
	}

	// Sort elves in order of their total
	sort.Sort(sort.Reverse(elves))

	fmt.Printf("Top elf total calories: %d\n", elves[0].Total())

	topThree := elves[0].Total() + elves[1].Total() + elves[2].Total()
	fmt.Printf("Top 3 elves total calories: %d\n", topThree)
}
