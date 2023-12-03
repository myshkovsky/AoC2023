package main

import (
	"AoC2023/solutions/day2/part1"
	"os"
)

func main() {
    var files [2]*os.File
    var err error
    files[0], err = os.Open("./input.txt")
    defer files[0].Close()
    if err != nil {
        panic(err)
    }
    part1.Solve(files[0])
}
