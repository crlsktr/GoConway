package main

import (
	"os"
)

var size int = 8

func main() {
	// generateFile()
	args := os.Args[1:]
	var inputFile, outputFileName string

	if len(args) > 1 {
		inputFile = args[0]
		outputFileName = args[1]
	} else {
		inputFile = "input"
		outputFileName = "output"
	}

	board := initializeBoard()
	board = loadData(inputFile, board)
	next := nextGeneration(board)

	writeData(outputFileName, next)
}
