package main

import (
	"testing"
)

func TestFloodFill(t *testing.T) {
	board := MakeBoard(5)
	snake := MakeSnake(Coord{1, 0}, Coord{1, 1}, Coord{1, 2}, Coord{0, 2}, Coord{0, 3}, Coord{0, 4})
	board.Snakes = append(board.Snakes, snake)

	ff := MakeFF(board)
	ff.FloodFill(Coord{0, 0})

	expected := 2

	if ff.count != expected {
		t.Errorf("Error floodfill exp: %v actual %v", expected, ff.count)
	}

	ff = MakeFF(board)
	ff.FloodFill(Coord{2, 0})

	expected = 18
	if ff.count != expected {
		t.Errorf("Error floodfill exp: %v actual %v", expected, ff.count)
	}
}

func TestFloodFillCount(t *testing.T) {
	board := MakeBoard(5)
	snake := MakeSnake(Coord{1, 0}, Coord{1, 1}, Coord{1, 2}, Coord{0, 2}, Coord{0, 3}, Coord{0, 4})
	board.Snakes = append(board.Snakes, snake)
	gameState = MakeGameState(board, snake)

	actual := floodFillCount(gameState, "left")

	expected := 2

	if actual != expected {
		t.Errorf("Error floodfill exp: %v actual %v", expected, actual)
	}

	actual = floodFillCount(gameState, "right")

	expected = 18

	if actual != expected {
		t.Errorf("Error floodfill exp: %v actual %v", expected, actual)
	}
}
