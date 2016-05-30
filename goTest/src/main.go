package main

import "fmt"

var points = [][]byte{
	{0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 1, 1, 0}}

var jobsCount int = 1;

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
	checkLeft_1 := func(p Position) Position {
		return Position{1,1}
	}
	checkRight_1 := func(p Position) Position {
		return Position{1,1}
	}
	checkTop_1 := func(p Position) Position {
		return Position{1,1}
	}
	checkBot_1 := func(p Position) Position {
		return Position{1,1}
	}
	return check(p,
		checkLeft_1,
		checkRight_1,
		checkTop_1,
		checkBot_1)
}
func check(p Position, f... func(p Position, ) Position) Position {

	var min Position = p;
	if p.x > 0 {
		min = checkMin(min, checkLeft(Position{p.x - 1, p.y}))
	}
	if p.x < len(points[0]) - 1 {
		min = checkMin(min, checkRight(Position{p.x + 1, p.y}))
	}
	if p.y > 0 {
		min = checkMin(min, checkTop(Position{p.x, p.y - 1}))
	}

	if p.y < len(points) - 1 {
		min = checkMin(min, checkBottom(Position{p.x, p.y + 1}))
	}
	return min
}
func checkLeft(p Position) Position {
	var min Position = p;
	fmt.Println("check left", p)
	if points[p.y][p.x] == 0 {
		return Position{p.x + 1, p.y}
	}
	if p.x > 0 {
		min = checkMin(min, checkLeft(Position{p.x - 1, p.y}))
	}
	if p.y > 0 {
		min = checkMin(min, checkTop(Position{p.x, p.y - 1}))
	}

	if p.y < len(points) - 1 {
		min = checkMin(min, checkBottom(Position{p.x, p.y + 1}))
	}
	fmt.Println("left return", min)
	return min
}
func checkRight(p Position) Position {
	var min Position = p;
	fmt.Println("check right", p)
	if points[p.y][p.x] == 0 {
		return Position{p.x - 1, p.y}
	}
	if p.x < len(points[0]) - 1 {
		min = checkMin(min, checkRight(Position{p.x + 1, p.y}))
	}
	if p.y > 0 {
		min = checkMin(min, checkTop(Position{p.x, p.y - 1}))
	}

	if p.y < len(points) - 1 {
		min = checkMin(min, checkBottom(Position{p.x, p.y + 1}))
	}
	fmt.Println("right return", min)
	return min
}
func checkTop(p Position) Position {
	var min Position = p;
	fmt.Println("check top", p)
	if points[p.y][p.x] == 0 {
		return Position{p.x, p.y + 1}
	}
	if p.x > 0 {
		min = checkMin(min, checkLeft(Position{p.x - 1, p.y}))
	}
	if p.x < len(points[0]) - 1 {
		min = checkMin(min, checkRight(Position{p.x + 1, p.y}))
	}
	if p.y > 0 {
		min = checkMin(min, checkTop(Position{p.x, p.y - 1}))
	}
	fmt.Println("top return", min)
	return min
}
func checkBottom(p Position) Position {
	var min Position = p;
	fmt.Println("check bot", p)
	if points[p.y][p.x] == 0 {
		return Position{p.x, p.y - 1}
	}
	if p.x > 0 {
		min = checkMin(min, checkLeft(Position{p.x - 1, p.y}))
	}
	if p.x < len(points[0]) - 1 {
		min = checkMin(min, checkRight(Position{p.x + 1, p.y}))
	}
	if p.y < len(points) - 1 {
		min = checkMin(min, checkBottom(Position{p.x, p.y + 1}))
	}
	fmt.Println("bot return", min)
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

type Position struct {
	x, y int;
}

