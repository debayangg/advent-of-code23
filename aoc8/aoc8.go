package aoc8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Aoc8() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	path := ""
	left_right := make(map[string][]string)
	start := []string{}

	for {
		line, err := reader.ReadString('\n')

		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		if strings.Contains(line, "=") {
			key := strings.Split(line, " = ")[0]
			value := strings.Split(line, " = ")[1]
			value = strings.Trim(value, "()")
			values := strings.Split(value, ", ")
			left_right[key] = values
			if key[len(key)-1] == 'A' {
				start = append(start, key)
			}

		} else {
			if line != "" {
				path = line
			}
		}

		if err != nil {
			break
		}
	}

	idx, steps := 0, 0
	step := []int{}

	for _, val := range start {
		str := val
		steps, idx = 0, 0
		for str[len(str)-1] != 'Z' {
			char := string(path[idx])
			if char == "R" {
				str = left_right[str][1]
			} else if char == "L" {
				str = left_right[str][0]
			}
			steps++
			idx = (idx + 1) % len(path)
		}
		step = append(step, steps)
	}

	lcm := 1
	for _, val := range step {
		lcm = LCM(lcm, val)
	}
	fmt.Println(lcm)

}
