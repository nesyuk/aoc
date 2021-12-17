package main

import (
	"fmt"
	"math"
)

const (
	//x1, x2 = 20, 30
	//y1, y2 = -10, -5
	x1, x2 = 155, 182
	y1, y2 = -117, -67
)

var (
	target = area{coord{x: x1, y: y1}, coord{x: x2, y: y2}}
)

type area struct {
	from coord
	to   coord
}

type coord struct {
	x, y int
}

func (c *coord) isWithin(a area) bool {
	return c.x >= a.from.x && c.x <= a.to.x && c.y >= a.from.y && c.y <= a.to.y
}

func (c *coord) isOutsideBounds(a area) bool {
	return c.x > a.to.x || c.y < a.from.y
}

func (c *coord) move(velocity coord) coord {
	c.x += velocity.x
	c.y += velocity.y

	next := velocity
	next.y -= 1
	if velocity.x > 0 {
		next.x -= 1
	} else if velocity.x < 0 {
		next.x += 1
	}
	return next
}

func launch(pos *coord, velocity coord) (maxHeight int, reached bool) {
	for {
		velocity = pos.move(velocity)
		if maxHeight < pos.y {
			maxHeight = pos.y
		}
		if pos.isWithin(target) {
			reached = true
			return
		}
		if pos.isOutsideBounds(target) {
			maxHeight = math.MinInt
			return
		}
	}
}

func solution() (int, int) {
	maxHeight, count := math.MinInt, 0
	for y := -200; y < 200; y++ {
		for x := -200; x < 200; x++ {
			height, reached := launch(&coord{0, 0}, coord{x, y})
			if maxHeight < height {
				maxHeight = height
			}
			if reached {
				count++
			}
		}
	}
	return maxHeight, count
}

func main() {
	fmt.Println("Day17")

	maxHeight, count := solution()
	fmt.Printf("max height: %d, count: %d\n", maxHeight, count)
}
