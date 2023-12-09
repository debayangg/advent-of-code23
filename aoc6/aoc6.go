package aoc6

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Aoc6() {
	input, _ := os.Open("input.txt")
	reader := bufio.NewReader(input)
	time_dist := make(map[string]int)

	for {
		line, err := reader.ReadString('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		split_vals := strings.Split(line, ":")
		split_vals[1] = strings.Trim(split_vals[1], " ")
		time_dist[split_vals[0]] = 0

		for _, val := range strings.Split(split_vals[1], " ") {
			if val != "" {
				num, _ := strconv.Atoi(val)
				if time_dist[split_vals[0]] == 0 {
					time_dist[split_vals[0]] = num
				} else {
					tens := int(math.Pow(10, float64(len(val))))
					time_dist[split_vals[0]] = time_dist[split_vals[0]]*tens + num
				}
			}
		}

		if err != nil {
			break
		}
	}
	fmt.Println(time_dist)
	start, end := 0, 0

	for i := 1; i < time_dist["Time"]; i++ {
		if (time_dist["Time"]-i)*i > time_dist["Distance"] && start == 0 {
			start = i
		}
		if (time_dist["Time"]-i)*i > time_dist["Distance"] {
			end = i
		}
	}
	/*
		for i := time_dist["Time"] - 1; i >= start; i-- {
			if (time_dist["Time"]-i)*i > time_dist["Distance"] {
				end = i
				break
			}
		}
	*/
	fmt.Println(start, end)
	fmt.Println(end - start + 1)
}
