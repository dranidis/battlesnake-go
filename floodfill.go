package main

type FF struct {
	board [][]bool
	count int
}

func MakeFF(board Board) FF {
	var ff FF
	ff.count = 0

	ff.board = make([][]bool, board.Height)
	for i := range ff.board {
		ff.board[i] = make([]bool, board.Width)
	}

	for i := 0; i < board.Height; i++ {
		for j := 0; j < board.Width; j++ {
			ff.board[i][j] = false
		}
	}

	for _, snake := range board.Snakes {
		// fill except for the tail
		for i := 0; i < len(snake.Body)-1; i++ {
			ff.board[snake.Body[i].X][snake.Body[i].Y] = true
		}
	}
	return ff
}

func (ff *FF) FloodFill(coord Coord) {
	if !ff.inside(coord) || ff.isSet(coord) {
		return
	}
	ff.set(coord)
	ff.FloodFill(Coord{coord.X, coord.Y + 1})
	ff.FloodFill(Coord{coord.X, coord.Y - 1})
	ff.FloodFill(Coord{coord.X + 1, coord.Y})
	ff.FloodFill(Coord{coord.X - 1, coord.Y})
}

func (ff FF) inside(coord Coord) bool {
	return coord.X >= 0 && coord.X < len(ff.board) && coord.Y >= 0 && coord.Y < len(ff.board)
}

func (ff FF) isSet(coord Coord) bool {
	return ff.board[coord.X][coord.Y]
}

func (ff *FF) set(coord Coord) {
	ff.board[coord.X][coord.Y] = true
	ff.count++
}

func floodFillCount(state GameState, move string) int {
	myHead := state.You.Head

	ff := MakeFF(state.Board)
	var co Coord
	switch move {
	case "up":
		co = Coord{myHead.X, myHead.Y + 1}
	case "down":
		co = Coord{myHead.X, myHead.Y - 1}
	case "right":
		co = Coord{myHead.X + 1, myHead.Y}
	case "left":
		co = Coord{myHead.X - 1, myHead.Y}
	}
	ff.FloodFill(co)

	return ff.count
}
