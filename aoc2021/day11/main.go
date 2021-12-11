package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	N = 10
	V = 9
	STEPS = 100
)

func flash(idxs [][]int8, octos *[][]int8) int {
	flashed := 0
	for len(idxs) > 0 {
		i, j := idxs[0][0], idxs[0][1]
		idxs = idxs[1:]

		if (*octos)[i][j] > V {
			flashed++
			(*octos)[i][j] = 0
			for x := int8(-1); x <= int8(1); x++ {
				for y := int8(-1); y <= int8(1); y++ {
					i1, j1 := i + x, j + y
					if i1 >= 0 && i1 < N && j1 >= 0 && j1 < N && (*octos)[i1][j1] != 0 {
						(*octos)[i1][j1]++
						if (*octos)[i1][j1] > V {
							idxs = append(idxs, []int8{i1, j1})
						}
					}
				}
			}
		}
	}
	return flashed
}

func step(octos *[][]int8) int {
	flashed := 0
	for i := int8(0); i < N; i++ {
		for j := int8(0); j < N; j++ {
			(*octos)[i][j]++
		}
	}
	for i := int8(0); i < N; i++ {
		for j := int8(0); j < N; j++ {
			if (*octos)[i][j] > V {
				flashed += flash([][]int8{{i, j}}, octos)
			}
		}
	}
	return flashed
}

func solution2(grid [][]int8) int {
	steps := 0
	for flashed := 0; flashed != N*N; steps++ {
		flashed = step(&grid)
	}
	return steps
}

func solution1(grid [][]int8, steps int) int {
	flashes := 0
	for s := 0; s < steps; s++ {
		flashed := step(&grid)
		flashes += flashed
	}
	return flashes
}

func main() {
	fmt.Println("Day11")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([][]int8, 0)
	for scanner.Scan() {
		line := make([]int8, 0)
		for _, str := range strings.Split(scanner.Text(), "") {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, int8(num))
		}
		lines = append(lines, line)
	}
	flashed := solution1(lines, STEPS)
	fmt.Printf("solution 1: flashed: %d\n", flashed)

	steps := solution2(lines)
	fmt.Printf("solution 2: steps: %d\n", steps + STEPS)
}
