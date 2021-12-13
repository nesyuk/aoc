package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dot struct {
	x, y int
}

const (
	X = "x"
	Y = "y"
)

type Instruction struct {
	direction string
	line      int
}

type Paper struct {
	dots map[Dot]bool
	h, w int
}

func NewPaper(dots [][]int) Paper {
	w, h := 0, 0
	paper := Paper{dots: map[Dot]bool{}}
	for _, dot := range dots {
		paper.dots[Dot{
			x: dot[0],
			y: dot[1],
		}] = true
		if dot[0] > w {
			w = dot[0]
		}
		if dot[1] > h {
			h = dot[1]
		}
	}
	paper.w = w
	paper.h = h
	return paper
}

func (p *Paper) print() {
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			if p.dots[Dot{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("================")
}

func (p *Paper) Fold(instruction Instruction) {
	switch instruction.direction {
	case Y:
		p.foldHoriz(instruction.line)
	case X:
		p.foldVert(instruction.line)
	default:
		log.Fatal("unknown instruction")
	}
}

func (p *Paper) foldHoriz(line int) {
	folded := map[Dot]bool{}
	for dot := range p.dots {
		if dot.y > line {
			folded[Dot{x: dot.x, y: line - (dot.y - line)}] = true
		} else {
			folded[dot] = true
		}
	}
	p.dots = folded
	p.h = line
}

func (p *Paper) foldVert(line int) {
	folded := map[Dot]bool{}
	for dot := range p.dots {
		if dot.x > line {
			folded[Dot{x: line - (dot.x - line), y: dot.y}] = true
		} else {
			folded[dot] = true
		}
	}
	p.dots = folded
	p.w = line
}

func solution2(paper Paper, instructions []Instruction) {
	for _, instruction := range instructions {
		paper.Fold(instruction)
	}
	fmt.Println("Code:")
	paper.print()
}

func solution1(paper Paper, instructions []Instruction) int {
	paper.Fold(instructions[0])
	return len(paper.dots)
}

func main() {
	fmt.Println("Day13")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dots := make([][]int, 0)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			break
		}
		coords := strings.Split(str, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}
		dots = append(dots, []int{x, y})
	}
	paper := NewPaper(dots)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		str := scanner.Text()
		strs := strings.Split(str, " ")
		instrStr := strings.Split(strs[2], "=")
		line, err := strconv.Atoi(instrStr[1])
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, Instruction{direction: instrStr[0], line: line})
	}

	totalDots := solution1(paper, instructions)
	fmt.Printf("solution 1: totalDots: %d\n", totalDots)

	solution2(paper, instructions)
}
