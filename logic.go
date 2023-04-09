package main

import (
	"log"
	"math/rand"
)

func getNextMove(state GameState) string {
	isMoveSafe := MakeBooleanMap(true)
	mayCollideWithLargerOrEqualHead := MakeBooleanMap(false)

	AvoidWall(state, isMoveSafe)
	AvoidAllSnakes(state, isMoveSafe)
	FindPossibleLosingHeadCollisions(state, mayCollideWithLargerOrEqualHead)

	totallySafeMoves, partiallySafeMoves := getNextSafeMoves(isMoveSafe, mayCollideWithLargerOrEqualHead)

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	nextMove := chooseNextMove(state, totallySafeMoves, partiallySafeMoves)
	return nextMove
}

func getNextSafeMoves(isMoveSafe map[string]bool, mayCollideWithLargerOrEqualHead map[string]bool) ([]string, []string) {
	totallySafeMoves := []string{}
	partiallySafeMoves := []string{}

	for move, isSafe := range isMoveSafe {
		if isSafe && !mayCollideWithLargerOrEqualHead[move] {
			totallySafeMoves = append(totallySafeMoves, move)
			partiallySafeMoves = append(partiallySafeMoves, move)
		} else if isSafe {
			partiallySafeMoves = append(partiallySafeMoves, move)
		}
	}
	log.Printf("Total Safe: %v", totallySafeMoves)
	log.Printf("Part  Safe: %v", partiallySafeMoves)
	return totallySafeMoves, partiallySafeMoves
}

func chooseNextMove(state GameState, totallysafeMoves []string, partiallySafeMoves []string) string {
	var nextMoves []string

	if len(totallysafeMoves) > 0 {
		nextMoves = totallysafeMoves
	} else if len(partiallySafeMoves) > 0 {
		nextMoves = partiallySafeMoves
	} else {
		nextMoves = []string{"down"}
	}

	nextMoves = MaxFloodFillMoves(state, nextMoves)

	return chooseAMoveTowardsFood(state, nextMoves)
}

func MaxFloodFillMoves(state GameState, nextMoves []string) []string {
	maxMoves := []string{}
	floodFillValues := make(map[string]int)

	for _, move := range nextMoves {
		floodFill := floodFillCount(state, move)
		floodFillValues[move] = floodFill
		if floodFill > len(state.You.Body) {
			maxMoves = append(maxMoves, move)
		}
	}
	log.Printf("FF %v\n", floodFillValues)

	if len(maxMoves) > 0 {
		return maxMoves
	} else {
		// select larger floodfill
		// todo: select aread with tail if tail is reachable
		max := 0
		maxKey := ""
		for key, value := range floodFillValues {
			if value > max {
				max = value
				maxKey = key
			}
		}
		return []string{maxKey}
	}
}

func chooseAMoveTowardsFood(state GameState, moves []string) string {
	foodMoves := []string{}
	moveTowardsFood := MakeBooleanMap(false)
	FindFood(state, moves, moveTowardsFood)
	for move, isFood := range moveTowardsFood {
		if isFood {
			foodMoves = append(foodMoves, move)
		}
	}
	if len(foodMoves) > 0 {
		return randomChoice(foodMoves)
	}
	return randomChoice(moves)
}

func randomChoice(moves []string) string {
	return moves[rand.Intn(len(moves))]
}

func FindPossibleLosingHeadCollisions(state GameState, mayCollideWithLargerOrEqualHead map[string]bool) {
	for _, snake := range state.Board.Snakes {
		if snake.ID != state.You.ID && snake.Length >= state.You.Length {
			findHeadCollisions(mayCollideWithLargerOrEqualHead, state, snake)
		}
	}
}

func AvoidAllSnakes(state GameState, isMoveSafe map[string]bool) {
	for _, snake := range state.Board.Snakes {
		avoidSnake(isMoveSafe, state, snake)
	}
}

func AvoidWall(state GameState, isMoveSafe map[string]bool) {
	myHead := state.You.Head // Coordinates of your head
	isMoveSafe["left"] = myHead.X != 0
	isMoveSafe["right"] = myHead.X != state.Board.Width-1
	isMoveSafe["down"] = myHead.Y != 0
	isMoveSafe["up"] = myHead.Y != state.Board.Height-1

}

func avoidSnake(isMoveSafe map[string]bool, state GameState, snake Battlesnake) {
	myHead := state.You.Head // Coordinates of your head

	for index, part := range snake.Body {
		if index == len(snake.Body)-1 {
			break
		}

		if myHead.Y == part.Y {
			if myHead.X == part.X+1 {
				isMoveSafe["left"] = false
			} else if myHead.X == part.X-1 {
				isMoveSafe["right"] = false
			}
		}
		if myHead.X == part.X {
			if myHead.Y == part.Y+1 {
				isMoveSafe["down"] = false
			} else if myHead.Y == part.Y-1 {
				isMoveSafe["up"] = false
			}
		}
	}
}

func findHeadCollisions(mayCollideWithHead map[string]bool, state GameState, snake Battlesnake) {
	myHead := state.You.Head // Coordinates of your head
	for _, otherNextHead := range nextMoves(snake.Head) {
		if myHead.Y == otherNextHead.Y {
			if myHead.X-1 == otherNextHead.X {
				mayCollideWithHead["left"] = true
			} else if myHead.X+1 == otherNextHead.X {
				mayCollideWithHead["right"] = true
			}
		}
		if myHead.X == otherNextHead.X {
			if myHead.Y-1 == otherNextHead.Y {
				mayCollideWithHead["down"] = true
			} else if myHead.Y+1 == otherNextHead.Y {
				mayCollideWithHead["up"] = true
			}
		}
	}
}

func nextMoves(snakeHead Coord) []Coord {
	return []Coord{{snakeHead.X + 1, snakeHead.Y}, {snakeHead.X - 1, snakeHead.Y}, {snakeHead.X, snakeHead.Y + 1}, {snakeHead.X, snakeHead.Y - 1}}
}

func MakeBooleanMap(value bool) map[string]bool {
	return map[string]bool{
		"up":    value,
		"down":  value,
		"left":  value,
		"right": value,
	}
}
