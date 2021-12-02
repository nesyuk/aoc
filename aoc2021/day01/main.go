package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solution2(data []int) int {
	counter := 0
	prev := data[0] + data[1] + data[2]
	for i := 3; i < len(data); i++ {
		cur := prev - data[i - 3] + data[i]
		if cur > prev {
			counter++
		}
		prev = cur
	}
	return counter
}

func solution1(data []int) int {
	counter := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println("Day1")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Failed to read integer: %s", str)
		}
		input = append(input, num)
	}
	fmt.Printf("solution 1: %d\n", solution1(input))
	fmt.Printf("solution 2: %d\n", solution2(input))
}
