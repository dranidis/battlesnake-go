package main

import (
	"errors"
)

func FindFood(state GameState, possibleMoves []string, moveTowardsFood map[string]bool) {
	myHead := state.You.Head // Coordinates of your head

	food, err := getClosestFood(state)
	if err != nil {
		return
	}

	for _, move := range possibleMoves {
		switch move {
		case "down":
			moveTowardsFood[move] = myHead.Y > food.Y
		case "up":
			moveTowardsFood[move] = myHead.Y < food.Y
		case "left":
			moveTowardsFood[move] = myHead.X > food.X
		case "right":
			moveTowardsFood[move] = myHead.X < food.X
		}
	}
}

func getClosestFood(state GameState) (Coord, error) {
	myHead := state.You.Head // Coordinates of your head

	if len(state.Board.Food) == 0 {
		return Coord{0, 0}, errors.New("NO FOOD")
	}

	var closestFood Coord
	closestDistance := state.Board.Height + state.Board.Width

	for _, food := range state.Board.Food {
		distanceToFood := distance(myHead, food)
		if distanceToFood < closestDistance {
			closestDistance = distanceToFood
			closestFood = food
		}
	}
	return closestFood, nil

}
