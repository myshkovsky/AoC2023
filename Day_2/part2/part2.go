package part2

import (
	"AoC2023/solutions/day2/part1"
	"bufio"
	"fmt"
	"os"
)

func Solve(file *os.File) {
    s := bufio.NewScanner(file)
    games := 0
    for s.Scan() {
        m := part1.GetHighestRGB(s)
        games += m["red"] * m["green"] * m["blue"]
    }
    fmt.Println("Total sum of set powers: ", games)
}
