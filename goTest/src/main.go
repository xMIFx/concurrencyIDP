package main

import "fmt"

type Position struct {
	x, y int;
}
type Checker func(checker Checker, previousPosition Position, currentPosition Position, nextPositions... Position) Position

var points = [][]byte{
	{0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{0, 1, 1, 0, 1}}

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
	/*var checked []Position = []Position{}*/
	check := func(checker Checker, previousPosition Position, currentPosition Position, nextPositions... Position) Position{
		if points[currentPosition.y][currentPosition.x] == 0 {
			return previousPosition
		}
		var min Position = currentPosition;
		for _, p := range nextPositions {
			if p.x < currentPosition.x && currentPosition.x > 0 {
				min = checkMin(min, checker(checker, currentPosition, p, getNextPositions(p, 1, 0)...))
			}
			if p.x > currentPosition.x && currentPosition.x < len(points[0]) - 1 {
				min = checkMin(min, checker(checker, currentPosition, p, getNextPositions(p, -1, 0)...))
			}
			if p.y < currentPosition.y && currentPosition.y > 0 {
				min = checkMin(min, checker(checker, currentPosition, p, getNextPositions(p, 0, 1)...))
			}
			if p.y > currentPosition.y && currentPosition.y < len(points) - 1 {
				min = checkMin(min, checker(checker, currentPosition, p, getNextPositions(p, 0, -1)...))
			}

		}
		return min
	}
	return check(check, p, p, getNextPositions(p, 0, 0)...)
}

func getNextPositions(p Position, excludeX int, excludeY int) ([]Position) {
	if excludeX == -1 {
		return []Position{Position{p.x + 1, p.y}, Position{p.x, p.y - 1}, Position{p.x, p.y + 1}}
	}
	if excludeX == 1 {
		return []Position{Position{p.x - 1, p.y}, Position{p.x, p.y - 1}, Position{p.x, p.y + 1}}
	}
	if excludeY == -1 {
		return []Position{Position{p.x + 1, p.y}, Position{p.x - 1, p.y}, Position{p.x, p.y + 1}}
	}
	if excludeY == 1 {
		return []Position{Position{p.x + 1, p.y}, Position{p.x - 1, p.y}, Position{p.x, p.y - 1}}
	}
	return []Position{Position{p.x + 1, p.y}, Position{p.x - 1, p.y}, Position{p.x, p.y - 1}, Position{p.x, p.y + 1}}
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
func equals(p Position, p1 Position) bool {

	return (p.y == p1.y &&  p.x == p1.x)
}
func contains(s []Position, e Position) bool {
	for _, a := range s {
		if a.x == e.x && a.y == e.y {
			return true
		}
	}
	return false
}


