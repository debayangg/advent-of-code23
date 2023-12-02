package aoc2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Aoc2() {
	reader := bufio.NewReader(os.Stdin)
	var sum int = 0
	for {
		var line string
		line, err := reader.ReadString('\n')

		reg1 := regexp.MustCompile("(?i)[0-9]+ [A-Za-z]+")
		balls := reg1.FindAllString(line, -1)

		vals := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, ball := range balls {
			number, _ := strconv.Atoi(strings.Split(ball, " ")[0])
			colour := strings.Split(ball, " ")[1]
			vals[colour] = int(math.Max(float64(number), float64(vals[colour])))
		}

		sum += vals["red"] * vals["green"] * vals["blue"]

		if err != nil {
			break
		}
	}

	fmt.Println(sum)
}
