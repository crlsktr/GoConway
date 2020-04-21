package main

import (
	"fmt"
	"os"
)

var size int = 8

func main() {
	// generateFile()
	args := os.Args[1:]
	var inputFile, outputFile string

	if len(args) > 1 {
		inputFile = args[0]
		outputFile = args[2]
	} else {
		inputFile = "input"
		outputFile = "output"
	}

	board := initializeBoard()
	board = loadData(inputFile, board)
	next := nextGeneration(board)

	writeData(outputFile, next)
	for _, row := range next {
		fmt.Println(row, ",")
	}
}
