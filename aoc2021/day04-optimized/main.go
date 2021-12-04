package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var WINNER, _ = strconv.ParseUint("11111", 2, 64)

func solution2(boards []*board, input []uint64) (uint64, uint64) {
	leftToWin := len(boards)
	for _, n := range input {
		for _, b := range boards {
			if b.hasWon() {
				continue
			}
			b.mark(n)
			if b.hasWon() {
				if leftToWin == 1 {
					return n, b.getSumOfUnmarked()
				}
				b.won = true
				leftToWin--
			}
		}
	}
	return 0, 0
}

func solution1(boards []*board, input []uint64) (uint64, uint64) {
	for _, n := range input {
		for _, b := range boards {
			b.mark(n)
			if b.hasWon() {
				return n, b.getSumOfUnmarked()
			}
		}
	}
	return 0, 0
}

type pos struct {
	row    uint64
	col    uint64
	marked bool
}

type board struct {
	rows  []uint64
	cols  []uint64
	index map[uint64]pos
	won   bool
}

func newboard() *board {
	return &board{rows: make([]uint64, 5), cols: make([]uint64, 5), index: map[uint64]pos{}}
}

func (b *board) mark(n uint64) {
	p, exist := b.index[n]
	if !exist {
		return
	}
	b.rows[p.row] |= 1 << (4 - p.col)
	b.cols[p.col] |= 1 << (4 - p.row)

	p.marked = true
	b.index[n] = p
}

func (b *board) hasWon() (won bool) {
	if b.won {
		return true
	}
	for i := uint64(0); i < 5; i++ {
		if b.rows[i] == WINNER || b.cols[i] == WINNER {
			return true
		}
	}
	return false
}

func (b *board) getSumOfUnmarked() uint64 {
	sum := uint64(0)
	for n, p := range b.index {
		if !p.marked {
			sum += n
		}
	}
	return sum
}

func main() {
	fmt.Println("Day4")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read numbers
	scanner.Scan()
	str := scanner.Text()
	numStrs := strings.Split(str, ",")

	input := make([]uint64, 0)
	for _, s := range numStrs {
		input = append(input, parseUint(s))
	}

	// read boards
	boards := make([]*board, 0)

	for scanner.Scan() {
		str = scanner.Text()
		b := newboard()
		s := uint64(0)
		for i := uint64(0); i < 5; i++ {
			numStrs = make([]string, 0)
			scanner.Scan()
			str = scanner.Text()
			splitted := strings.Split(str, " ")
			for _, n := range splitted {
				if n != "" {
					numStrs = append(numStrs, n)
				}
			}
			for j := uint64(0); j < 5; j++ {
				n := parseUint(numStrs[j])
				b.index[n] = pos{i, j, false}
				s += n
			}
		}
		boards = append(boards, b)
	}
	//num, sum := solution1(boards, input)
	//fmt.Printf("solution 1: num: %d, sum: %d, ans: %d\n", num, sum, num*sum)

	num, sum := solution2(boards, input)
	fmt.Printf("solution 1: num: %d, sum: %d, ans: %d\n", num, sum, num*sum)
}

func parseUint(s string) uint64 {
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
