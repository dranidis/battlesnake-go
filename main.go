package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"log"
	"math/rand"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "DD",       // TODO: Your Battlesnake username
		Color:      "#0033cc",  // TODO: Choose color
		Head:       "caffeine", // TODO: Choose head
		Tail:       "coffee",   // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {
	log.Printf("TURN %d: %d %d\n", state.Turn, state.You.Head.X, state.You.Head.Y)

	isMoveSafe := MakeTrueMap()

	surelyNotCollidesWithHead := MakeTrueMap()

	AvoidWall(state, isMoveSafe)

	AvoidAllSnakes(state, isMoveSafe)

	AvoidHeadCollisions(state, surelyNotCollidesWithHead)

	// Are there any safe moves left?
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

	nextMove := ""

	if len(totallysafeMoves) > 0 {
		nextMove = totallysafeMoves[rand.Intn(len(totallysafeMoves))]
	} else if len(partiallySafeMoves) > 0 {
		nextMove = partiallySafeMoves[rand.Intn(len(partiallySafeMoves))]
	} else {
		nextMove = "down"
	}

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)
	return BattlesnakeMoveResponse{Move: nextMove}
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

func main() {
	RunServer()
}
