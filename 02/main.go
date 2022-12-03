package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	RockPoints     int = 1
	PaperPoints    int = 2
	ScissorsPoints int = 3
)

const (
	LosePoints int = 0
	DrawPoints int = 3
	WinPoints  int = 6
)

var (
	roundOneResults map[byte]int = map[byte]int{
		0x5: RockPoints + DrawPoints,
		0x9: RockPoints + LosePoints,
		0xd: RockPoints + WinPoints,
		0x6: PaperPoints + WinPoints,
		0xa: PaperPoints + DrawPoints,
		0xe: PaperPoints + LosePoints,
		0x7: ScissorsPoints + LosePoints,
		0xb: ScissorsPoints + WinPoints,
		0xf: ScissorsPoints + DrawPoints,
	}

	roundTwoResults map[byte]int = map[byte]int{
		0x6: RockPoints + DrawPoints,
		0x9: RockPoints + LosePoints,
		0xf: RockPoints + WinPoints,
		0xa: PaperPoints + DrawPoints,
		0xd: PaperPoints + LosePoints,
		0x7: PaperPoints + WinPoints,
		0xe: ScissorsPoints + DrawPoints,
		0x5: ScissorsPoints + LosePoints,
		0xb: ScissorsPoints + WinPoints,
	}
)

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

	var totalScore int
	for _, round := range data {
		score := roundOneResults[round]
		totalScore += score
	}
	fmt.Printf("Round One Score: %d\n", totalScore)

	totalScore = 0
	for _, round := range data {
		score := roundTwoResults[round]
		totalScore += score
	}
	fmt.Printf("Round Two Score: %d\n", totalScore)
}

func parseData(r io.Reader) ([]byte, error) {
	var out []byte

	s := bufio.NewScanner(r)

	for s.Scan() {
		b := s.Bytes()
		x, y := (b[0] - 64), (b[2] - 87)
		out = append(out, y|x<<2)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
