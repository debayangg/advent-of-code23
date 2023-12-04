package aoc4

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func Aoc4() {
	sum := 0
	var card_match []int
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		line = line[:len(line)-1]
		var correct, test []string

		for i, val := range strings.Split(line, "|") {
			if i == 0 {
				idx := strings.Index(val, ":")
				correct = strings.Split(strings.Trim(val[idx+1:], " "), " ")
			} else {
				test = strings.Split(strings.Trim(val, " "), " ")
			}
		}

		matches := 0

		for i, val := range test {
			if val == "" {
				test = append(test[:i], test[i+1:]...)
			}
		}

		for i, val := range correct {
			if val == "" {
				correct = append(correct[:i], correct[i+1:]...)
			}
		}

		for _, val := range correct {
			if slices.Contains(test, val) {
				matches++
			}
		}
		card_match = append(card_match, matches)
		/*if matches > 0 {
			sum += int(math.Pow(2.0, float64(matches-1)))
		}*/

		if err != nil {
			break
		}
	}

	card := make(map[int]int)

	for i := range card_match {
		card[i+1] = 1
	}

	for i, match := range card_match {
		sum += card[i+1]
		fmt.Println(i+1, match, card[i+1])
		for j := 1; j <= match; j++ {
			_, check := card[j+i+1]
			if check {
				card[i+j+1] += card[i+1]
			}
		}
	}

	fmt.Println(sum)
}
