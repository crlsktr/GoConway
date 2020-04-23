package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func loadData(fileName string, board [][]byte) [][]byte {
	data, _ := ioutil.ReadFile(fileName)
	fmt.Println("loaded data from file: ", fileName)
	if float64(len(data)) == math.Pow(float64(size), 2) {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				board[i][j] = data[(i*size)+j]
			}
		}
	} else {
		println("data loaded is invalid or has the wrong size")
	}

	for _, row := range board {
		fmt.Println(row)
	}
	return board
}

func writeData(fileName string, board [][]byte) {
	fmt.Println("write data to file: ", fileName)
	f, err := os.Create(fileName)
	check(err)

	defer f.Close()

	for _, row := range board {
		fmt.Println(row)
		f.Write(row)
	}
	f.Sync()
}
