package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func filter(data []uint64, bit uint64, need uint64) []uint64 {
	sorted := map[uint64][]uint64{0: {}, 1: {}}

	for _, num := range data {
		if flag := num & (1 << bit); flag > 0 {
			sorted[0] = append(sorted[0], num)
		} else {
			sorted[1] = append(sorted[1], num)
		}
	}
	mostCommon := uint64(0)
	if  len(sorted[0]) >= len(sorted[1]) {
		mostCommon = 1
	}
	return sorted[need ^ mostCommon]
}

func solution2(data []uint64, n uint64) (uint64, uint64) {
	filtered := data
	for i := uint64(0); i < n && len(filtered) > 1; i++ {
		filtered = filter(filtered, n - i - 1, 1)
	}
	oxygen := filtered[0]

	filtered = data
	for i := uint64(0); i < n && len(filtered) > 1; i++ {
		filtered = filter(filtered, n - i - 1, 0)
	}
	co2 := filtered[0]

	return oxygen, co2
}

func solution1(data []uint64, n uint64) (uint64, uint64) {
	counts := make([]uint64, n)
	for _, num := range data {
		for i := uint64(0); i < n; i++ {
			if flag := num & (1 << i); flag > 0 {
				counts[n-i-1] += 1
			}
		}
	}
	half := uint64(len(data) / 2)
	gamma := uint64(0)
	epsilon := uint64(0)
	for i := range counts {
		if counts[n-uint64(i)-1] > half {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	return gamma, epsilon
}

func main() {
	fmt.Println("Day2")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]uint64, 0)
	scanner := bufio.NewScanner(file)
	size := 0
	for scanner.Scan() {
		str := scanner.Text()
		size = len(str)
		num, err := strconv.ParseUint(str, 2, 64)
		if err != nil {
			log.Fatalf("Failed to read integer: %s", str)
		}
		input = append(input, num)
	}
	gamma, epsilon := solution1(input, uint64(size))
	fmt.Printf("solution 1: gamma: %d, epsilon: %d, ans: %d\n", gamma, epsilon, gamma*epsilon)

	oxygen, co2 := solution2(input, uint64(size))
	fmt.Printf("solution 2: oxygen: %d, co2: %d, ans: %d\n", oxygen, co2, oxygen*co2)
}
