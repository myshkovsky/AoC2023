package main

import (
	solution "AoC2023/solutions/day6/parts"
	"os"
)

func main() {
	path := "./input.txt"
	var files [2]*os.File
	var err error
	files[0], err = os.Open(path)
	files[1], err = os.Open(path)
	defer files[0].Close()
	defer files[1].Close()
	if err != nil {
		panic(err)
	}
	solution.Part1(files[0])
}
