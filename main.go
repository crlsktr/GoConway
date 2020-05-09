package main

import (
	"fmt"
	"os"
)

var size int = 2000

func main() {
	args := os.Args[1:]
	var inputFile, outputFileName string

	if len(args) ==2 {
		inputFile = args[0]
		outputFileName = args[1]
	} else if len(args) == 1 && args[0] == "generate"{
		generateFile()
		return
	} else {
		fmt.Println("Found no arguments, please provide arguments [inputFile] [outputFile]")
		// inputFile = "input"
		// outputFileName = "output"
		return
	}

	board := initializeBoard()
	board = loadData(inputFile, board)
	next := nextGeneration(board)

	writeData(outputFileName, next)
}
