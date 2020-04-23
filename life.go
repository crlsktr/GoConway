package main

func initializeBoard() [][]byte {
	board := make([][]byte, size)
	for i := 0; i < size; i++ {
		board[i] = make([]byte, size)
	}
	return board
}

func nextGeneration(board [][]byte) [][]byte {
	newBoard := initializeBoard()
	for ri, row := range board {
		for ci, el := range row {
			neighborsCount := countNeighbors(board, ri, ci)
			newBoard[ri][ci] = cellLiveOrDie(neighborsCount, el)
		}
	}
	return newBoard
}

func countNeighbors(board [][]byte, rowIndex int, columnIndex int) int {
	count := 0
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if !(i == 0 && j == 0) &&
				(rowIndex+i >= 0 && rowIndex+i < len(board)) &&
				(columnIndex+j >= 0 && columnIndex+j < len(board)) {
				if board[rowIndex+i][columnIndex+j] == 1 {
					count++
				}
			}
		}
	}
	return count
}

func cellLiveOrDie(count int, cell byte) byte {
	var val byte
	if cell == 1 {
		if count < 2 || count > 3 {
			val = 0
		} else {
			val = 1
		}
	} else if cell == 0 {
		if count == 3 {
			val = 1
		} else {
			val = 0
		}
	}
	return val
}
