package main

import (
	"testing"
)

func TestClosestWithNoFood(t *testing.T) {
	board := MakeBoard(5)
	state := MakeGameState(board, MakeSnake(Coord{1, 1}))

	_, err := getClosestFood(state)

	if err == nil {
		t.Errorf("Error should be thrown")
	}
}

func TestClosestWithOneFood(t *testing.T) {
	board := AddFood(MakeBoard(5), Coord{4, 4})
	state := MakeGameState(board, MakeSnake(Coord{1, 1}))

	food, err := getClosestFood(state)

	if err != nil {
		t.Errorf("Error should not be thrown")
	}
	if food.X != 4 && food.Y != 4 {
		t.Errorf("food not found")
	}
}

func TestClosestWithTwoFood(t *testing.T) {
	board := AddFood(MakeBoard(5), Coord{4, 4}, Coord{0, 0})
	state := MakeGameState(board, MakeSnake(Coord{1, 1}))

	food, err := getClosestFood(state)

	if err != nil {
		t.Errorf("Error should not be thrown")
	}
	if food.X != 0 && food.Y != 0 {
		t.Errorf("closest food not found")
	}
}
