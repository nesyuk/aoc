package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func (c *coord) moveHorizontally(other *coord) []coord {
	x1, x2 := c.x, other.x
	coords := make([]coord, 0)
	if x1 < x2 {
		for ; x1 <= x2; x1++ {
			coords = append(coords, coord{x: x1, y: c.y})
		}
	} else {
		for ; x1 >= x2; x1-- {
			coords = append(coords, coord{x: x1, y: c.y})
		}
	}
	return coords
}

func (c *coord) moveVertically(other *coord) []coord {
	y1, y2 := c.y, other.y
	coords := make([]coord, 0)

	if y1 < y2 {
		for ; y1 <= y2; y1++ {
			coords = append(coords, coord{x: c.x, y: y1})
		}
	} else {
		for ; y1 >= y2; y1-- {
			coords = append(coords, coord{x: c.x, y: y1})
		}
	}
	return coords
}

func (c *coord) String() string {
	return fmt.Sprintf("x: %d, y: %d", c.x, c.y)
}

func (c *coord) parse(s string) {
	strs := strings.Split(s, ",")
	x, err := strconv.Atoi(strs[0])
	if err != nil {
		log.Fatal(err)
	}
	c.x = x
	y, err := strconv.Atoi(strs[1])
	if err != nil {
		log.Fatal(err)
	}
	c.y = y
}

type move struct {
	from *coord
	to *coord
}

func (m *move) isHorizontal() bool {
	return m.from.y == m.to.y
}

func (m *move) isVertical() bool {
	return m.from.x == m.to.x
}

func (m *move) String() string {
	return fmt.Sprintf("from: %d, to: %d", m.from, m.to)
}

func (m *move) parse(s string) {
	strs := strings.Split(s, " -> ")
	from, to := new(coord), new(coord)
	from.parse(strs[0])
	to.parse(strs[1])
	m.from = from
	m.to = to
}

func solution(moves []*move) int {
	points := map[coord]int{}
	for _, m := range moves {
		coords := make([]coord, 0)
		switch {
		case m.isHorizontal():
			coords = m.from.moveHorizontally(m.to)
		case m.isVertical():
			coords = m.from.moveVertically(m.to)
		default: // for solution 2:
			coordsX := m.from.moveHorizontally(m.to)
			coordsY := m.from.moveVertically(m.to)
			for i := range coordsX {
				coords = append(coords, coord{x: coordsX[i].x, y: coordsY[i].y})
			}
		}
		for _, c := range coords {
			points[c] += 1
		}
	}
	count := 0
	for _, p := range points {
		if p >= 2 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Day5")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moves := make([]*move, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := new(move)
		m.parse(scanner.Text())
		moves = append(moves, m)
	}
	//count := solution(moves)
	//fmt.Printf("solution 1: ans: %d\n", count)

	count := solution(moves)
	fmt.Printf("solution 2: ans: %d\n", count)
}
