package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var MEMO = map[int]map[string]map[string]uint64{} // times to mutate -> sequence -> new chars count
var SEEN = map[string]uint64{}

func mutate(l, r *string, rules *map[string]string, time int, times *int) map[string]uint64 {
	acc := map[string]uint64{}

	lr := fmt.Sprintf("%s%s", *l, *r)
	if timesMemo, exist := MEMO[*times-time]; exist {
		if lrMemo, exist := timesMemo[lr]; exist {
			for m, count := range lrMemo {
				SEEN[m] += count
			}
			return lrMemo
		}
	}
	if time > *times {
		return acc
	}
	m := (*rules)[lr]
	SEEN[m]++

	acc[m] = 1
	for k, v := range mutate(l, &m, rules, time+1, times) {
		acc[k] += v
	}
	for k, v := range mutate(&m, r, rules,time+1, times) {
		acc[k] += v
	}

	if _, exist := MEMO[*times-time]; !exist {
		MEMO[*times-time] = make(map[string]map[string]uint64, 0)
	}
	timesMemo := MEMO[*times-time]
	if _, exist := timesMemo[lr]; !exist {
		timesMemo[lr] = make(map[string]uint64, 0)
	}
	timesMemo[lr] = acc

	return acc
}

func solution1(sequence string, rules *map[string]string, times int) uint64 {
	for _, ch := range sequence {
		SEEN[string(ch)]++
	}
	for i := 1; i < len(sequence); i++ {
		l, r := string(sequence[i-1]), string(sequence[i])
		mutate(&l, &r, rules, 1, &times)
	}
	minCount, maxCount := uint64(math.MaxUint64), uint64(0)
	for _, count := range SEEN {
		if minCount > count {
			minCount = count
		}
		if maxCount < count {
			maxCount = count
		}
	}
	return maxCount - minCount
}

func main() {
	fmt.Println("Day14")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sequence := scanner.Text()

	rules := make(map[string]string, 0)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			continue
		}
		ruleMap := strings.Split(str, " -> ")
		rules[ruleMap[0]] = ruleMap[1]
	}

	result := solution1(sequence, &rules, 40)
	fmt.Printf("solution 1: result: %d\n", result)
}
