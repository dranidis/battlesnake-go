package main

import (
	"testing"
)

var gameState GameState
var isMoveSafe map[string]bool
var mayCollideWithLargerHead map[string]bool
var moveTowardsFood map[string]bool

func beforeEach() {
	isMoveSafe = MakeBooleanMap(true)
	mayCollideWithLargerHead = MakeBooleanMap(false)
	moveTowardsFood = MakeBooleanMap(false)
	gameState = MakeGameState(MakeBoard(11), Battlesnake{})
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

// Head collisions

func TestAvoidHeadColisionWithEqual(t *testing.T) {
	beforeEach()
	givenABoardOfSize(5)
	givenYourSnakeBodyIs(Coord{2, 2}, Coord{2, 1}, Coord{2, 0})
	givenAnotherSnakeBodyIs(Coord{3, 3}, Coord{4, 3}, Coord{4, 4})
	whenFindPossibleHeadColisions()
	thenMoveMayCollideWithLargerHead(t, "right")
	thenMoveMayCollideWithLargerHead(t, "up")
}

func TestDoNotAvoidHeadColisionWithSmaller(t *testing.T) {
	beforeEach()
	givenABoardOfSize(5)
	givenYourSnakeBodyIs(Coord{2, 2}, Coord{2, 1}, Coord{2, 0})
	givenAnotherSnakeBodyIs(Coord{3, 3}, Coord{4, 3})
	whenFindPossibleHeadColisions()
	thenMoveDoesNotCollideWithLargerHead(t, "right")
	thenMoveDoesNotCollideWithLargerHead(t, "up")
}

// Food

func TestMoveTowardsClosestFood(t *testing.T) {
	beforeEach()
	givenABoardOfSize(5)
	givenFoodAt(Coord{4, 4})
	givenYourSnakeBodyIs(Coord{2, 2}, Coord{2, 1}, Coord{2, 0})
	whenMoveTowardsClosestFood()
	thenMoveIsTowardsFood(t, "right")
	thenMoveIsTowardsFood(t, "up")
}

func thenMoveIsTowardsFood(t *testing.T, s string) {
	if !moveTowardsFood[s] {
		t.Errorf("Move %v is not towards food at head %v", s, gameState.You.Head)
	}
}

func whenMoveTowardsClosestFood() {
	possibleMoves := []string{"up", "down", "left", "right"}
	FindFood(gameState, possibleMoves, moveTowardsFood)
}

func givenFoodAt(coord ...Coord) {
	gameState.Board.Food = append(gameState.Board.Food, coord...)
}

// Implementation of BDD functions

func thenMoveMayCollideWithLargerHead(t *testing.T, s string) {
	if !mayCollideWithLargerHead[s] {
		t.Errorf("Move %v does not collide at head %v", s, gameState.You.Head)
	}
}

func thenMoveDoesNotCollideWithLargerHead(t *testing.T, s string) {
	if mayCollideWithLargerHead[s] {
		t.Errorf("Move %v may collide at head %v", s, gameState.You.Head)
	}
}

func whenFindPossibleHeadColisions() {
	FindPossibleLosingHeadCollisions(gameState, mayCollideWithLargerHead)
}

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
	gameState.You = MakeSnake(coord...)
	gameState.You.ID = "me"
	gameState.Board.Snakes = append(gameState.Board.Snakes, gameState.You)
}

func givenAnotherSnakeBodyIs(coord ...Coord) {
	other := MakeSnake(coord...)
	other.ID = "other"
	gameState.Board.Snakes = append(gameState.Board.Snakes, other)
}

func givenABoardOfSize(size int) {
	gameState.Board = MakeBoard(size)
}
