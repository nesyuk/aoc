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

func parseNum(str string) uint64 {
	num, err := strconv.ParseUint(str, 2, len(str))
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func parseVersionPackId(encoded *string) (version uint64, packId uint64) {
	version = parseNum((*encoded)[:3])
	*encoded = (*encoded)[3:]

	packId = parseNum((*encoded)[:3])
	*encoded = (*encoded)[3:]

	return
}

func parseLiteralValue(encoded *string) uint64 {
	numStr := ""
	for {
		group := (*encoded)[:5]
		*encoded = (*encoded)[5:]
		numStr += group[1:]

		if group[0] == '0' {
			break
		}
	}
	num := parseNum(numStr)
	return num
}

func parse(encoded *string, versions *uint64) (result uint64) {
	if len(*encoded) == 0 || (len(*encoded) < 11 && parseNum(*encoded) == 0) {
		return
	}

	version, packId := parseVersionPackId(encoded)
	*versions += version

	// literal value
	if packId == 4 {
		return parseLiteralValue(encoded)
	}

	// operators
	i := (*encoded)[0]
	*encoded = (*encoded)[1:]

	nums := make([]uint64, 0)

	switch i {
	case '0':
		totLen := int(parseNum((*encoded)[:15]))
		*encoded = (*encoded)[15:]
		subpackages := (*encoded)[:totLen]
		*encoded = (*encoded)[totLen:]
		for len(subpackages) > 0 {
			nums = append(nums, parse(&subpackages, versions))
		}

	case '1':
		totCount := int(parseNum((*encoded)[:11]))
		*encoded = (*encoded)[11:]
		for n := 0; n < totCount; n++ {
			nums = append(nums, parse(encoded, versions))
		}
	}
	switch packId {
	case 0: // sum
		for _, n := range nums {
			result += n
		}
	case 1: // product
		result = 1
		for _, n := range nums {
			result *= n
		}
	case 2: // min
		result = math.MaxInt
		for _, n := range nums {
			if result > n {
				result = n
			}
		}
	case 3: // max
		for _, n := range nums {
			if result < n {
				result = n
			}
		}
	case 5: // gt
		if nums[0] > nums[1] {
			result = 1
		}
	case 6: // lt
		if nums[0] < nums[1] {
			result = 1
		}
	case 7:
		if nums[0] == nums[1] {
			result = 1
		}
	default:
		log.Fatalf("unknown package id: %v", packId)
	}
	return result
}

func solution1(binary string) (uint64, uint64) {
	versions := uint64(0)
	return parse(&binary, &versions), versions
}

func toBinary(hex string) string {
	str := strings.Split(hex, "")
	binary := make([]string, 0)
	for _, ch := range str {
		i, err := strconv.ParseUint(ch, 16, 16)
		if err != nil {
			log.Fatal(err)
		}
		binary = append(binary, fmt.Sprintf("%04b", i))
	}
	return strings.Join(binary, "")
}

func main() {
	fmt.Println("Day16")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	result, versionsSum := solution1(toBinary(scanner.Text()))
	fmt.Println(result)
	fmt.Println(versionsSum)
}
