package aoc1

import "fmt"

func match(sub string, check int) int {
	var f, l int
	if len(sub) >= 3 {

		if check == 1 {
			f, l = len(sub)-3, len(sub)
		} else {
			f, l = 0, 3
		}

		switch sub[f:l] {

		case "one":
			return 1
		case "two":
			return 2
		case "six":
			return 6

		}
	}

	if len(sub) >= 4 {

		if check == 1 {
			f, l = len(sub)-4, len(sub)
		} else {
			f, l = 0, 4
		}

		switch sub[f:l] {

		case "four":
			return 4
		case "five":
			return 5
		case "nine":
			return 9
		}
	}

	if len(sub) >= 5 {

		if check == 1 {
			f, l = len(sub)-5, len(sub)
		} else {
			f, l = 0, 5
		}
		switch sub[f:l] {

		case "three":
			return 3
		case "seven":
			return 7
		case "eight":
			return 8
		}
	}

	return -1
}

func Aoc1() {
	sum := 0
	for {
		var line string
		_, err := fmt.Scanln(&line)

		if err != nil {
			break
		}

		var last, first int = -1, -1

		for i := 0; i < len(line); i++ {
			if (line[i]-'0') <= 9 && (line[i]-'0') >= 0 {
				first = int(line[i] - '0')
				break
			}
			val := match(line[0:i+1], 1)
			if val != -1 {
				first = val
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if (line[i]-'0') <= 9 && (line[i]-'0') >= 0 {
				last = int(line[i] - '0')
				break
			}

			val := match(line[i:], 0)

			if val != -1 {
				last = val
				break
			}
		}
		sum += first*10 + last
	}

	fmt.Println(sum)
}
