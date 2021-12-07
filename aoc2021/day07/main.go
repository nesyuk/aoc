package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func solution2(crabs map[int]int, minPos, maxPos int) (int, int) {
	minDistance, minCrabPos := math.MaxInt, -1
	fuelMemo := map[int]int{}

	for pos := minPos; pos <= maxPos; pos++ {
		distance := 0
		for crabPos, crabCount := range crabs {
			diff := absDiff(pos, crabPos)
			fuel, exist := fuelMemo[diff]
			if !exist {
				for i := 1; i <= diff; i++ {
					fuel += i
				}
				fuelMemo[diff] = fuel
			}
			distance += fuel * crabCount
		}
		if minDistance > distance {
			minDistance = distance
			minCrabPos = pos
		}
	}
	return minCrabPos, minDistance
}

func solution1(crabs map[int]int, minPos, maxPos int) (int, int) {
	minDistance, minCrabPos := math.MaxInt, -1

	for pos := minPos; pos <= maxPos; pos++ {
		distance := 0
		for crabPos, crabCount := range crabs {
			distance += absDiff(pos, crabPos) * crabCount
		}
		if minDistance > distance {
			minDistance = distance
			minCrabPos = pos
		}
	}
	return minCrabPos, minDistance
}

func absDiff(v1, v2 int) int {
	if v1 > v2 {
		return v1 - v2
	}
	return v2 - v1
}

func main() {
	fmt.Println("Day7")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), ",")

	minPos, maxPos := math.MaxInt, 0
	crabs := make(map[int]int, 0)
	for _, s := range strs {
		pos, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		crabs[pos]++
		if minPos > pos {
			minPos = pos
		}
		if maxPos < pos {
			maxPos = pos
		}
	}

	minPos, minDistance := solution1(crabs, minPos, maxPos)
	fmt.Printf("solution 1: position: %d, distance: %d ans: %d\n", minPos, minDistance, minDistance)

	minPos, minDistance = solution2(crabs, minPos, maxPos)
	fmt.Printf("solution 2: position: %d, distance: %d ans: %d\n", minPos, minDistance, minDistance)
}
