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
	row, col int
}

func walk(floor [][]int8, stack []coord) []coord {
	x, y := len(floor[0]), len(floor)
	seen := map[coord]bool{}
	for len(stack) > 0 {
		c := stack[0]
		stack = stack[1:]
		seen[c] = true

		for _, n := range []coord{
			{c.row - 1, c.col},
			{c.row + 1, c.col},
			{c.row, c.col - 1},
			{c.row, c.col + 1}} {

			if n.col < 0 || n.col > x-1 || n.row < 0 || n.row > y-1 || floor[n.row][n.col] == 9 {
				continue
			}
			if !seen[n] {
				stack = append(stack, n)
			}
		}
	}
	basin := make([]coord, 0)
	for c := range seen {
		basin = append(basin, c)
	}
	return basin
}

func solution2(floor [][]int8, coords []coord) (int, int, int) {
	l1, l2, l3 := 0, 0, 0
	for _, c := range coords {
		basin := walk(floor, []coord{c})
		l := len(basin)
		if l1 < l {
			l3, l2, l1 = l2, l1, l
		} else if l2 < l {
			l3, l2 = l2, l
		} else if l3 < l {
			l3 = l
		}
	}
	return l1, l2, l3
}

func solution1(floor [][]int8) (int64, []coord) {
	count := int64(0)
	coords := make([]coord, 0)
	x, y := len(floor[0]), len(floor)
	for col := 0; col < x; col++ {
	ROW:
		for row := 0; row < y; row++ {
			center := floor[row][col]

			for _, n := range []coord{
				{row - 1, col},
				{row + 1, col},
				{row, col - 1},
				{row, col + 1}} {

				if n.col < 0 || n.col > x-1 || n.row < 0 || n.row > y-1 {
					continue
				}
				if floor[row][col] >= floor[n.row][n.col] {
					continue ROW
				}
			}
			count += int64(center) + 1
			coords = append(coords, coord{row, col})
		}
	}
	return count, coords
}

func main() {
	fmt.Println("Day8")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	floor := make([][]int8, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "")
		heights := make([]int8, 0)
		for _, str := range strs {
			height, err := strconv.ParseInt(str, 10, 8)
			if err != nil {
				log.Fatal(err)
			}
			heights = append(heights, int8(height))
		}
		floor = append(floor, heights)
	}
	sumOfLows, coords := solution1(floor)
	fmt.Printf("solution 1: sumOfLows: %d\n", sumOfLows)

	l1, l2, l3 := solution2(floor, coords)
	fmt.Printf("solution 2: l1: %d, l2: %d l3: %d, prod: %d \n", l1, l2, l3, l1*l2*l3)
}
