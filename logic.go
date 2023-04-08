package main

import (
	"log"
	"math/rand"
)

func getNextMove(state GameState) string {
	isMoveSafe := MakeTrueMap()

	surelyNotCollidesWithHead := MakeTrueMap()

	AvoidWall(state, isMoveSafe)

	AvoidAllSnakes(state, isMoveSafe)

	AvoidHeadCollisions(state, surelyNotCollidesWithHead)
	totallysafeMoves := []string{}
	partiallySafeMoves := []string{}

	for move, isSafe := range isMoveSafe {
		if isSafe && surelyNotCollidesWithHead[move] {
			totallysafeMoves = append(totallysafeMoves, move)
			partiallySafeMoves = append(partiallySafeMoves, move)
		} else if isSafe {
			partiallySafeMoves = append(partiallySafeMoves, move)
		}
	}
	log.Printf("Total Safe: %v", totallysafeMoves)
	log.Printf("Part  Safe: %v", partiallySafeMoves)

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	nextMove := ""

	if len(totallysafeMoves) > 0 {
		nextMove = totallysafeMoves[rand.Intn(len(totallysafeMoves))]
	} else if len(partiallySafeMoves) > 0 {
		nextMove = partiallySafeMoves[rand.Intn(len(partiallySafeMoves))]
	} else {
		nextMove = "down"
	}
	return nextMove
}

func AvoidHeadCollisions(state GameState, surelyNotCollidesWithHead map[string]bool) {
	for _, snake := range state.Board.Snakes {
		if snake.ID != state.You.ID && snake.Length >= state.You.Length {
			avoidHead(surelyNotCollidesWithHead, state, snake)
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

func avoidHead(isMoveSafe map[string]bool, state GameState, snake Battlesnake) {
	myHead := state.You.Head // Coordinates of your head
	for _, otherNextHead := range nextMoves(snake.Head) {
		if myHead.Y == otherNextHead.Y {
			if myHead.X-1 == otherNextHead.X {
				isMoveSafe["left"] = false
			} else if myHead.X+1 == otherNextHead.X {
				isMoveSafe["right"] = false
			}
		}
		if myHead.X == otherNextHead.X {
			if myHead.Y-1 == otherNextHead.Y {
				isMoveSafe["down"] = false
			} else if myHead.Y+1 == otherNextHead.Y {
				isMoveSafe["up"] = false
			}
		}
	}
}

func nextMoves(snakeHead Coord) []Coord {
	return []Coord{{snakeHead.X + 1, snakeHead.Y}, {snakeHead.X - 1, snakeHead.Y}, {snakeHead.X, snakeHead.Y + 1}, {snakeHead.X, snakeHead.Y - 1}}
}

func MakeTrueMap() map[string]bool {
	return map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}
}
