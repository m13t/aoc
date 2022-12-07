package main

import (
	"fmt"
	"os"

	"github.com/m13t/aoc/07/types"
)

const (
	FilesystemSize      int = 70000000
	UpdateSpaceRequired int = 30000000
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tree, err := types.Parse(f)
	if err != nil {
		panic(err)
	}

	total, err := roundOne(tree)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round One Total: %d\n", total)

	total, err = roundTwo(tree)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Round Two Total: %d\n", total)
}

func roundOne(tree *types.Entry) (int, error) {
	var out int

	tree.Walk(func(e *types.Entry) {
		if e.IsDir() && e.Size < 100000 {
			out += e.Size
		}
	})

	return out, nil
}

func roundTwo(tree *types.Entry) (int, error) {
	candidate := tree
	free := FilesystemSize - tree.Size

	tree.Walk(func(e *types.Entry) {
		if e.IsDir() && e.Size+free > UpdateSpaceRequired {
			if e.Size < candidate.Size {
				candidate = e
			}
		}
	})

	return candidate.Size, nil
}
