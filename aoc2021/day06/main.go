package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func nextDay(current []int) []int {
	next := make([]int, len(current))
	for i := 1; i < len(current); i++ {
		next[i-1] = current[i]
	}
	next[8] = current[0]
	next[6] += current[0]
	return next
}

func solution1(fish []int, days int) int {
	for i := 0; i < days; i++ {
		fish = nextDay(fish)
	}
	total := 0
	for _, count := range fish {
		total += count
	}
	return total
}

func main() {
	fmt.Println("Day6")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), ",")

	fish := make([]int, 9)
	for _, s := range strs {
		days, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fish[days]++
	}

	count := solution1(fish, 80)
	fmt.Printf("solution 1: ans: %d\n", count)

	count = solution1(fish, 256)
	fmt.Printf("solution 2: ans: %d\n", count)
}
