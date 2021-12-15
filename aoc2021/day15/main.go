package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row, col int
	prev     *coord
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func weight(row, col int, weights *[][]int) int {
	h, w := len(*weights), len((*weights)[0])
	v := (*weights)[row % h][col % w]
	scale := (row / h) + (col / w)
	return 1 + (v + scale - 1) % 9
}

func relax(v *Item, u *Item, pq *PriorityQueue, weights *[][]int) {
	c := v.value.(coord)
	w := weight(c.row, c.col, weights)
	if v.priority > u.priority + w {
		prev := u.value.(coord)
		c.prev = &prev
		pq.update(v, c, u.priority + w)
	}
}

func buildQueue(h, w int) (*PriorityQueue, [][]*Item){
	coords := make([][]*Item, h)
	for i, _ := range coords {
		coords[i] = make([]*Item, w)
	}

	pq := make(PriorityQueue, h*w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			priority := math.MaxInt
			c := coord{row: i, col: j, prev: nil}
			if i == 0 && j == 0 {
				priority = 0
			}
			idx := i*w + j
			coords[i][j] = &Item{value: c, priority: priority, index: idx}
			pq[idx] = coords[i][j]
		}
	}
	heap.Init(&pq)
	return &pq, coords
}

func dijkstra(weights *[][]int, h, w int) int {
	pq, coords := buildQueue(h, w)
	seen := map[coord]int{}

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Item)
		c := u.value.(coord)

		seen[c] = u.priority
		for _, n := range [][]int{{c.row, c.col-1}, {c.row, c.col+1}, {c.row-1, c.col}, {c.row+1, c.col}} {
			if n[0] >= 0 && n[0] < w &&n[1] >= 0 && n[1] < h {
				relax(coords[n[0]][n[1]], u, pq, weights)
			}
		}
	}
	c := coords[h-1][w-1].value.(coord)
	return seen[c]
}

func solution2(paths [][]int, scale int) int {
	h, w := len(paths), len(paths[0])
	return dijkstra(&paths, h*scale, w*scale)
}

func solution1(paths [][]int) int {
	h, w := len(paths), len(paths[0])
	coords := make([][]*Item, h)
	for i, _ := range coords {
		coords[i] = make([]*Item, w)
	}
	return dijkstra(&paths, h, w)
}

func main() {
	fmt.Println("Day15")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	paths := make([][]int, 0)
	for scanner.Scan() {
		rowStr := scanner.Text()
		colStrs := strings.Split(rowStr, "")
		row := make([]int, 0)
		for _, col := range colStrs {
			weight, err := strconv.Atoi(col)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, weight)
		}
		paths = append(paths, row)
	}
	shortest := solution1(paths)
	fmt.Printf("solution 1: result: %d\n", shortest)

	shortest = solution2(paths, 5)
	fmt.Printf("solution 2: result: %d\n", shortest)
}
