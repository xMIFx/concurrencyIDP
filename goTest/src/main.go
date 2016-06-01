package main

import "fmt"

type Position struct {
	x, y int;
}
type Checker func(checker Checker, currentPosition Position)

var points = [][]byte{
	{0, 1, 1, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 1, 1, 0},
	{0, 1, 1, 1, 0},
	{0, 1, 1, 1, 0},
	{0, 1, 1, 1, 0},
	{0, 1, 1, 1, 1}}

var jobsCount int = 8;

func main() {
	task := make(chan Position)
	done := make(chan bool)
	islandTop := make(chan Position)
	islands := make(chan []Position)
	go storeIslands(islandTop, islands);
	for i := 0; i < jobsCount; i++ {
		go startCountIsland(task, islandTop, done)
	}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points[i]); j++ {
			if points[i][j] == 1 {
				task <- Position{j, i}
			}
		}
	}
	for i := 0; i < jobsCount; i++ {
		task <- Position{-1, -1}
		<-done
	}

	islandTop <- Position{-1, -1}
	fmt.Println(<-islands)
}

func startCountIsland(tasks <- chan Position, islandsTop chan <- Position, done chan <- bool) {
	for {
		task := <-tasks
		if (task.x == -1) {
			break;
		}
		islandTop := checkPosition(task);
		islandsTop <- islandTop
	}
	done <- true
}

func storeIslands(islandsTop <- chan Position, islands chan <-[]Position) {
	var allTops []Position
	for {
		islandTop := <-islandsTop
		if islandTop.x == -1 {
			break
		}
		if !contains(allTops, islandTop) {
			allTops = append(allTops, islandTop)
		}
	}
	islands <- allTops
}

func checkPosition(p Position) Position {
	if p.x == 0 && p.y == 0 {
		return p;
	}
	var checked []Position = []Position{}
	var min Position = p;
	check := func(checker Checker, currentPosition Position) {
		if points[currentPosition.y][currentPosition.x] == 0 || contains(checked, currentPosition) {
			return
		}
		checked = append(checked, currentPosition)
		min = checkMin(min, currentPosition)

		if currentPosition.x > 0 {
			checker(checker, Position{currentPosition.x - 1, currentPosition.y})
		}
		if currentPosition.x < len(points[0]) - 1 {
			checker(checker, Position{currentPosition.x + 1, currentPosition.y})
		}
		if currentPosition.y > 0 {
			checker(checker, Position{currentPosition.x, currentPosition.y - 1})
		}
		if currentPosition.y < len(points) - 1 {
			checker(checker, Position{currentPosition.x, currentPosition.y + 1})
		}
	}
	check(check, p)
	return min
}

func checkMin(p Position, p1 Position) Position {
	if p.x == -1 {
		return p1
	}
	if p1.x == -1 {
		return p
	}

	if p.y < p1.y || (p.y == p1.y && p.x <= p1.x) {
		return p
	}
	return p1
}

func contains(s []Position, e Position) bool {
	for _, a := range s {
		if a.x == e.x && a.y == e.y {
			return true
		}
	}
	return false
}


