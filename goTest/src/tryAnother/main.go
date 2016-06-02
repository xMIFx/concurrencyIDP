package main

import (
	"fmt"
)

type Checker func(checker Checker, y int, x int)

var points = [][]byte{
	{1, 0, 0, 0, 1},
	{1, 0, 0, 1, 1},
	{1, 1, 1, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 0, 0, 1, 0},
	{0, 1, 0, 0, 0},
	{0, 1, 0, 1, 1}}

func main() {
	fmt.Println(points)
	var y, x int
	var count int = 0
	for {
		y, x = getPointWithOne(y, x)
		if y == -1 {
			break
		}
		replace(y, x)
		count++
	}
	fmt.Println(count)
}

func getPointWithOne(y int, x int) (int, int) {
	for i := y; i < len(points); i++ {
		for j := x; j < len(points[i]); j++ {
			if points[i][j] == 1 {
				return i, j
			}
		}
	}
	return -1, -1
}

func replace(y int, x int) {
	done := make(chan bool)
	countWorker := 0
	if x > 0 {
		countWorker++
		go replaceTask(y, x - 1, done)
	}
	if x < len(points[0]) - 1 {
		countWorker++
		go replaceTask(y, x + 1, done)
	}
	if y > 0 {
		countWorker++
		go replaceTask(y - 1, x, done)
	}
	if y < len(points) - 1 {
		countWorker++
		go replaceTask(y + 1, x, done)
	}
	for i := 0; i < countWorker; i++ {
		<-done
	}
}

func replaceTask(y int, x int, done chan <- bool) {

	var task = func(task Checker, y int, x int) {
		if points[y][x] == 0 {
			return
		}
		points[y][x] = 0
		if x > 0 {
			task(task, y, x - 1)
		}
		if x < len(points[0]) - 1 {
			task(task, y, x + 1)
		}
		if y > 0 {
			task(task, y - 1, x)
		}
		if y < len(points) - 1 {
			task(task, y + 1, x)
		}
	}
	task(task, y, x)
	done <- true
}
