package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	value     int
}

func solution2(commands []command) (int, int) {
	x, y, z := 0, 0, 0
	for _, comm := range commands {
		switch comm.direction {
		case "forward":
			x += comm.value
			y += comm.value * z
		case "up":
			z -= comm.value
		case "down":
			z += comm.value
		default:
			log.Fatalf("unknown direction: %s", comm.direction)
		}
	}
	return x, y
}

func solution1(commands []command) (int, int) {
	x, y := 0, 0
	for _, comm := range commands {
		switch comm.direction {
		case "forward":
			x += comm.value
		case "up":
			y -= comm.value
		case "down":
			y += comm.value
		default:
			log.Fatalf("unknown direction: %s", comm.direction)
		}
	}
	return x, y
}

func main() {
	fmt.Println("Day2")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]command, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		res := strings.Split(str, " ")
		num, err := strconv.Atoi(res[1])
		if err != nil {
			log.Fatalf("Failed to read integer: %s", str)
		}
		input = append(input, command{res[0], num})
	}
	x, y := solution1(input)
	fmt.Printf("solution 1: x: %d, y: %d, ans: %d\n", x, y, x*y)

	x, y = solution2(input)
	fmt.Printf("solution 2: x: %d, y: %d, ans: %d\n", x, y, x*y)
}
