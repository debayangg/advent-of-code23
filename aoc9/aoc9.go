package aoc9

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func sum(num []int) int {
	sum := 0
	for i := range num {
		val := num[i] * int(math.Pow(-1, float64(i)))
		sum += val
	}
	return sum
}

func all_zero(num []int) bool {
	for i := range num {
		val := num[i]
		if val != 0 {
			return false
		}
	}
	return true
}

func value(line []int) int {
	copy := line
	first_vals := make([]int, 0)
	for !all_zero(copy) {
		first_vals = append(first_vals, copy[0])
		new_line := make([]int, 0)
		prev := 0
		for i := range copy {
			val1 := copy[i]
			if i == 0 {
				prev = val1
				continue
			}

			new_line = append(new_line, val1-prev)
			prev = val1
		}
		copy = new_line
	}
	return sum(first_vals)

}

func Aoc9() {
	logfile, _ := os.Create("log.txt")
	log.SetOutput(logfile)

	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	input := make([][]int, 0)

	for {
		line, err := reader.ReadString('\n')

		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		line = strings.Trim(line, " ")
		int_vals := make([]int, 0)
		for _, str := range strings.Split(line, " ") {
			val, _ := strconv.Atoi(str)
			int_vals = append(int_vals, val)
		}
		input = append(input, int_vals)

		if err != nil {
			break
		}
	}

	sum := 0
	for _, str := range input {
		sum += value(str)
	}

	fmt.Println(sum)
}
