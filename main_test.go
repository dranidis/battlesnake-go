package main

import (
	"testing"
)

// type Board struct {
// 	Height  int           `json:"height"`
// 	Width   int           `json:"width"`
// 	Food    []Coord       `json:"food"`
// 	Hazards []Coord       `json:"hazards"`
// 	Snakes  []Battlesnake `json:"snakes"`
// }

var board Board

// type GameState struct {
// 	Game  Game        `json:"game"`
// 	Turn  int         `json:"turn"`
// 	Board Board       `json:"board"`
// 	You   Battlesnake `json:"you"`
// }

var gameState GameState

var isMoveSafe = MakeTrueMap()

func beforeEach() {
	isMoveSafe = MakeTrueMap()
}

// Avoid walls
func TestAvoidWallBottomRightCorner(t *testing.T) {
	beforeEach()
	givenABoardOfSize(3)
	givenYourSnakeBodyIs(Coord{0, 0})
	whenAvoidWalls()
	thenMoveIsNotSafe(t, "down")
	thenMoveIsNotSafe(t, "left")
}

func TestAvoidWallBottomLeftCorner(t *testing.T) {
	beforeEach()
	givenABoardOfSize(3)
	givenYourSnakeBodyIs(Coord{2, 0})
	whenAvoidWalls()
	thenMoveIsNotSafe(t, "down")
	thenMoveIsNotSafe(t, "right")
}

func TestAvoidWallTopLeftCorner(t *testing.T) {
	beforeEach()
	givenABoardOfSize(3)
	givenYourSnakeBodyIs(Coord{2, 2})
	whenAvoidWalls()
	thenMoveIsNotSafe(t, "up")
	thenMoveIsNotSafe(t, "right")
}

func TestAvoidWallTopRightCorner(t *testing.T) {
	beforeEach()
	givenABoardOfSize(3)
	givenYourSnakeBodyIs(Coord{0, 2})
	whenAvoidWalls()
	thenMoveIsNotSafe(t, "up")
	thenMoveIsNotSafe(t, "left")
}

// Avoid other snakes and yourself

func TestAvoidOtherSnake(t *testing.T) {
	beforeEach()
	givenABoardOfSize(5)
	givenYourSnakeBodyIs(Coord{1, 1})
	givenAnotherSnakeBodyIs(Coord{2, 2}, Coord{2, 1}, Coord{2, 0})
	whenAvoidSnakes()
	thenMoveIsNotSafe(t, "right")
	thenMoveIsSafe(t, "up")
	thenMoveIsSafe(t, "down")
	thenMoveIsSafe(t, "left")
}

func TestAvoidYourself(t *testing.T) {
	beforeEach()
	givenABoardOfSize(5)
	givenYourSnakeBodyIs(Coord{1, 1}, Coord{2, 1}, Coord{2, 2}, Coord{1, 2}, Coord{0, 2})
	whenAvoidSnakes()
	thenMoveIsNotSafe(t, "right") // collides with neck
	thenMoveIsNotSafe(t, "up")
	thenMoveIsSafe(t, "down")
	thenMoveIsSafe(t, "left")
}

func TestChaseYourTail(t *testing.T) {
	beforeEach()
	givenABoardOfSize(5)
	givenYourSnakeBodyIs(Coord{1, 1}, Coord{2, 1}, Coord{2, 2}, Coord{1, 2})
	whenAvoidSnakes()
	thenMoveIsNotSafe(t, "right") // collides with neck
	thenMoveIsSafe(t, "up")       // go to tail
	thenMoveIsSafe(t, "down")
	thenMoveIsSafe(t, "left")
}

// Implementation of BDD functions

func whenAvoidWalls() {
	AvoidWall(gameState, isMoveSafe)
}

func whenAvoidSnakes() {
	AvoidAllSnakes(gameState, isMoveSafe)
}

func thenMoveIsNotSafe(t *testing.T, s string) {
	if isMoveSafe[s] {
		t.Errorf("Move %v should not be safe at head %v", s, gameState.You.Head)
	}
}

func thenMoveIsSafe(t *testing.T, s string) {
	if !isMoveSafe[s] {
		t.Errorf("Move %v should be safe at head %v", s, gameState.You.Head)
	}
}

func givenYourSnakeBodyIs(coord ...Coord) {
	gameState.You = makeSnake(coord...)
	gameState.You.ID = "me"
	gameState.Board.Snakes = append(gameState.Board.Snakes, gameState.You)
}

func givenAnotherSnakeBodyIs(coord ...Coord) {
	other := makeSnake(coord...)
	other.ID = "other"
	gameState.Board.Snakes = append(gameState.Board.Snakes, other)
}

func makeSnake(coord ...Coord) Battlesnake {
	if len(coord) == 0 {
		panic("makeSnake: empty body")
	}
	return Battlesnake{"", "", 100, coord, coord[0], len(coord), "", "", Customizations{}}
}

func givenABoardOfSize(i int) {
	board = Board{i, i, []Coord{}, []Coord{}, []Battlesnake{}}
	gameState = GameState{Game{}, 0, board, Battlesnake{}}
}
