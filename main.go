package main

var size = 8

func initializeBoard() [][]bool {
	board := make([][]bool, size)
	for i := 0; i < size; i++ {
		board[i] = make([]bool, size)
	}
	return board
}

func main() {
	board := initializeBoard()
	// fmt.Println("before board:", board)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = true
		}
	}
}
