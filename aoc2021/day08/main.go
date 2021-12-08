package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solution2(displays [][][]int8) int {
	counter := 0
	for _, display := range displays {
		digitMap := map[int8]int{}
		signals, digits := display[0], display[1]
		var one, seven, four, five int8
		for _, signal := range signals {
			switch getLen(signal) {
			case 2:
				digitMap[signal] = 1
				one = signal
			case 3:
				digitMap[signal] = 7
				seven = signal
			case 4:
				digitMap[signal] = 4
				four = signal
			case 7:
				digitMap[signal] = 8
			}
		}
		for _, signal := range signals {
			l := getLen(signal)
			switch {
			case l == 5 && includes(signal, seven):
				digitMap[signal] = 3
			case l == 5 && union(signal, four) == 3:
				digitMap[signal] = 5
				five = signal
			case l == 5:
				digitMap[signal] = 2
			}
		}

		for _, signal := range signals {
			l := getLen(signal)
			switch {
			case l == 6 && includes(signal, five) && includes(signal, one):
				digitMap[signal] = 9
			case l == 6 && includes(signal, five):
				digitMap[signal] = 6
			case l == 6:
				digitMap[signal] = 0
			}
		}
		num := 0
		for i, digit := range digits {
			if i > 0 {
				num *= 10
			}
			num += digitMap[digit]
		}
		counter += num
	}
	return counter
}

func solution1(displays [][][]int8) int {
	counter := 0
	for _, display := range displays {
		for _, digit := range display[1] {
			l := getLen(digit)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				counter++
			}
		}
	}
	return counter
}

func getLen(digit int8) int {
	c := 0
	for i := 0; i < 8; i++ {
		if (digit & (1 << i)) > 0 {
			c++
		}
	}
	return c
}

func union(a, b int8) int {
	return getLen(a & b)
}

// checks a includes b
func includes(a, b int8) bool {
	return a&b == b
}

func toInt8(digit string) int8 {
	sorted := int8(0)
	for _, d := range digit {
		sorted |= 1 << (d - 'a')
	}
	return sorted
}

func main() {
	fmt.Println("Day8")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	displays := make([][][]int8, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " | ")
		signals := make([]int8, 0)
		for _, signal := range strings.Split(strs[0], " ") {
			signals = append(signals, toInt8(signal))
		}
		digits := make([]int8, 0)
		for _, digit := range strings.Split(strs[1], " ") {
			digits = append(digits, toInt8(digit))
		}
		displays = append(displays, [][]int8{signals, digits})
	}
	occurrences := solution1(displays)
	fmt.Printf("solution 1: occurences: %d\n", occurrences)

	numbersSum := solution2(displays)
	fmt.Printf("solution 2: numbers sum: %d\n", numbersSum)
}
