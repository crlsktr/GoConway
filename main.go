package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var size int = 2000

func main() {
	args := os.Args[1:]
	var inputFile, outputFileName string
	generations := 1

	if len(args) == 3 {
		inputFile = args[0]
		outputFileName = args[1]
		generationArg, err := strconv.Atoi(args[2])
		if err == nil{
			generations = generationArg
		}	
	} else if len(args) == 1 && args[0] == "generate"{
		generateFile()
		return
	} else {
		fmt.Println("Found no arguments, please provide arguments [inputFile] [outputFile] [numberOfGenerations]")
		return
	}

	
	var next [][]byte
	board := initializeBoard()
	board = loadData(inputFile, board)
	start := time.Now()
	for a:=0; a<generations; a++ {
		board = next
		next = nextGeneration(board)
	}

	duration := time.Since(start)
	fmt.Println("Computation of ",generations," generations lasted: ",duration)

	writeData(outputFileName, next)
}
