package main

import (
	"os"
)

func check(e error){
	if e!= nil {
		panic(e)
	}
}

func generateFile() {
	board := [][]byte{
		[]byte {0,0,0,0,0,0,0,0},
		[]byte {0,0,0,0,0,0,0,0},
		[]byte {0,0,0,0,0,0,0,0},
		[]byte {0,0,0,0,0,1,0,0},
		[]byte {0,0,0,1,0,1,0,0},
		[]byte {0,0,0,0,1,1,0,0},
		[]byte {0,0,0,0,0,0,0,0},
		[]byte {0,0,0,0,0,0,0,0},
	}

	f,err := os.Create("file")
	check(err)

	defer f.Close()

	for _,row := range board {
		f.Write(row)
	}

	f.Sync()
}