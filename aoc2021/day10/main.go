package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var (
	complement = map[rune]rune{'[': ']', '{': '}', '<': '>', '(': ')'}
)

func solution2(lines []string) int {
	totalPenaltiesIncomplete := make([]int, 0)
	penalty := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
LINES:
	for _, line := range lines {
		stack := make([]rune, 0)
		for _, bracket := range line {
			switch bracket {
			case '(', '[', '{', '<':
				stack = append([]rune{complement[bracket]}, stack...)
			case ']', '}', '>', ')':
				if len(stack) == 0 || stack[0] != bracket {
					// corrupted
					continue LINES
				} else {
					stack = stack[1:]
				}
			default:
				log.Fatalf("unknown character: %v", bracket)
			}
		}
		// incomplete
		counter := 0
		for len(stack) > 0 {
			counter *= 5
			counter += penalty[stack[0]]
			stack = stack[1:]
		}
		totalPenaltiesIncomplete = append(totalPenaltiesIncomplete, counter)
	}
	sort.Slice(totalPenaltiesIncomplete, func(i, j int) bool {
		return totalPenaltiesIncomplete[i] < totalPenaltiesIncomplete[j]
	})

	return totalPenaltiesIncomplete[len(totalPenaltiesIncomplete)/2]
}

func solution1(lines []string) int {
	totalPenaltyCorrupted := 0
	penalty := map[rune]int{']': 57, '}': 1197, '>': 25137, ')': 3}
LINES:
	for _, line := range lines {
		stack := make([]rune, 0)
		for _, bracket := range line {
			switch bracket {
			case '(', '[', '{', '<':
				stack = append([]rune{complement[bracket]}, stack...)
			case ']', '}', '>', ')':
				if len(stack) == 0 || stack[0] != bracket {
					// corrupted
					totalPenaltyCorrupted += penalty[bracket]
					continue LINES
				} else {
					stack = stack[1:]
				}
			default:
				log.Fatalf("unknown character: %v", bracket)
			}
		}
	}
	return totalPenaltyCorrupted
}

func main() {
	fmt.Println("Day10")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	penaltyCorrupted := solution1(lines)
	fmt.Printf("solution 1: penaltyCorrupted: %d\n", penaltyCorrupted)

	penaltyIncomplete := solution2(lines)
	fmt.Printf("solution 2: penaltyIncomplete: %d\n", penaltyIncomplete)
}
