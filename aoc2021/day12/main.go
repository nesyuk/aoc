package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	START = "start"
	END   = "end"
)

func isSmallCave(cave string) bool {
	return cave != START && cave != END && cave[0] >= 'a' && cave[0] <= 'z'
}

func visit(caves *map[string][]string, path []string, visited map[string]bool, visitTwice bool, paths *[][]string){
	last := path[len(path)-1]
	if last == END {
		*paths = append(*paths, path)
		return
	}
	for _, cave := range (*caves)[last] {
		if cave == START {
			continue
		}
		newVisited := make(map[string]bool)
		for c, v := range visited {
			newVisited[c] = v
		}
		newPath := append(path, cave)

		newVisitTwice := visitTwice
		if isSmallCave(cave) {
			if newVisited[cave] {
				if !visitTwice {
					continue
				}
				newVisitTwice = false
			}
			newVisited[cave] = true
		}
		visit(caves, newPath, newVisited, newVisitTwice, paths)
	}
}

func solution1(caves map[string][]string) int {
	allPaths := make([][]string, 0)
	for _, cave := range caves[START] {
		visited := map[string]bool{}
		if isSmallCave(cave) {
			visited[cave] = true
		}
		visit(&caves, []string{START, cave}, visited, false,  &allPaths)
	}
	return len(allPaths)
}

func solution2(caves map[string][]string) int {
	allPaths := make([][]string, 0)
	for _, cave := range caves[START] {
		visited := map[string]bool{}
		if isSmallCave(cave) {
			visited[cave] = true
		}
		visit(&caves, []string{START, cave}, visited, true,  &allPaths)
	}
	return len(allPaths)
}

func main() {
	fmt.Println("Day12")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	paths := map[string][]string{}
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "-")
		if _, exist := paths[strs[0]]; !exist {
			paths[strs[0]] = make([]string, 0)
		}
		paths[strs[0]] = append(paths[strs[0]], strs[1])

		// reverse path
		if _, exist := paths[strs[1]]; !exist {
			paths[strs[1]] = make([]string, 0)
		}
		paths[strs[1]] = append(paths[strs[1]], strs[0])
	}
	totalPaths := solution1(paths)
	fmt.Printf("solution 1: totalPaths: %d\n", totalPaths)

	totalPaths = solution2(paths)
	fmt.Printf("solution 2: totalPaths: %d\n", totalPaths)
}
