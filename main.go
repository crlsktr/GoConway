package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var size = 8

func initializeBoard() [][]byte {
	board := make([][]byte, size)
	for i := 0; i < size; i++ {
		board[i] = make([]byte, size)
	}
	return board
}

func main() {
	// generateFile()
	args := os.Args[1:]
	board := initializeBoard()
	data,_ := ioutil.ReadFile(args[0])

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i][j] = data[(i*size)+j]
		}
	}

	for _,row := range board {
		fmt.Println(row,",")
	}
}
