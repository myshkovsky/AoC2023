package solution

import (
	"bufio"
	"fmt"
	"os"
)

func getCardWinCount(winning map[int]bool, yours []int) int {
	wins := 0
	for _, v := range yours {
		if winning[v] {
			wins++
		}
	}
	return wins
}

func Part2(file *os.File) {
	s := bufio.NewScanner(file)
	cardId := 1
	cards := make(map[int]int)
	for s.Scan() {
		cards[cardId] += 1
		m := GetWinningNumbers(s.Text())
		nums := GetYourNumbers(s.Text())
		winCount := getCardWinCount(m, nums)
		for i := cards[cardId]; i > 0; i-- {
			for i := 1; i <= winCount; i++ {
				cards[cardId+i] += 1
			}
		}
		cardId++
	}
	sum := 0
    for i := cardId; i > 0; i-- {
		sum += cards[i]
	}
	fmt.Println("Part 2 solution: ", sum)
}
