package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inputFormat = "%d,%d -> %d,%d"

type coord struct {
	x, y int
}

func (c *coord) move(other *coord) (path []coord) {
	x1, x2, y1, y2 := c.x, other.x, c.y, other.y
	for x1 != x2 || y1 != y2 {
		path = append(path, coord{
			x: x1,
			y: y1,
		})
		if x1 > x2 {
			x1--
		} else if x1 < x2 {
			x1++
		}

		if y1 > y2 {
			y1--
		} else if y1 < y2 {
			y1++
		}
	}
	path = append(path, *other)
	return path
}

type move struct {
	from *coord
	to   *coord
}

func (m *move) isHorizontal() bool {
	return m.from.y == m.to.y
}

func (m *move) isVertical() bool {
	return m.from.x == m.to.x
}

func countPoints(points map[coord]int) int {
	count := 0
	for _, p := range points {
		if p >= 2 {
			count++
		}
	}
	return count
}

func solution2(moves []*move) int {
	points := map[coord]int{}
	for _, m := range moves {
		coords := m.from.move(m.to)
		for _, c := range coords {
			points[c] += 1
		}
	}
	return countPoints(points)
}

func solution1(moves []*move) int {
	points := map[coord]int{}
	for _, m := range moves {
		if m.isHorizontal() || m.isVertical() {
			coords := m.from.move(m.to)
			for _, c := range coords {
				points[c] += 1
			}
		}
	}
	return countPoints(points)
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
		x1, y1, x2, y2 := 0, 0, 0, 0
		if _, err = fmt.Sscanf(scanner.Text(), inputFormat, &x1, &y1, &x2, &y2); err != nil {
			log.Fatal(err)
		}
		moves = append(moves, &move{
			from: &coord{x: x1, y: y1},
			to:   &coord{x: x2, y: y2},
		})
	}
	count := solution1(moves)
	fmt.Printf("solution 1: ans: %d\n", count)

	count = solution2(moves)
	fmt.Printf("solution 2: ans: %d\n", count)
}
