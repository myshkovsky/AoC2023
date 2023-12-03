package main

import (
	"AoC2023/solutions/day1/part1"
	"AoC2023/solutions/day1/part2"
	"os"
)

func main() {
    var file [2]*os.File
    var err error
    file[0], err = os.Open("./input.txt")
    file[1], err = os.Open("./input.txt")
    if err != nil {
        panic(err)
    }
    defer file[0].Close()
    defer file[1].Close()
    part1.Solve(file[0])
    part2.Solve(file[1])
}
