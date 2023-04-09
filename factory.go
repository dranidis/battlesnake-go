package main

func MakeSnake(coord ...Coord) Battlesnake {
	if len(coord) == 0 {
		panic("makeSnake: empty body")
	}
	return Battlesnake{"", "", 100, coord, coord[0], len(coord), "", "", Customizations{}}
}

func MakeBoard(size int) Board {
	return Board{size, size, []Coord{}, []Coord{}, []Battlesnake{}}
}

func MakeGameState(board Board, snake Battlesnake) GameState {
	return GameState{Game{}, 0, board, snake}
}

func AddFood(board Board, coord ...Coord) Board {
	board.Food = append(board.Food, coord...)
	return board
}
