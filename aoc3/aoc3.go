package aoc3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type nums struct {
	start, end, line, val int
}

type pair struct {
	line, pos int
}

func Aoc3() {
	var vals []nums
	sym := make(map[pair]string)
	read := bufio.NewReader(os.Stdin)
	line_no := 0

	for {
		line, err := read.ReadString('\n')

		line = line[:len(line)-1]
		prev := ""
		index := 0

		for i := range line {
			if line[i]-'0' >= 0 && line[i]-'0' <= 9 {
				prev = prev + string(line[i])
			} else {

				if prev != "" {
					var value nums
					value.line = line_no
					value.start = index
					value.end = i - 1
					value.val, _ = strconv.Atoi(prev)
					index = i
					prev = ""
					vals = append(vals, value)
				}
				if line[i] == '.' {
					index++
				} else {
					index++
					sym[pair{line: line_no, pos: i}] = string(line[i])
				}
			}
		}

		if prev != "" {
			var value nums
			value.line = line_no
			value.start = index
			value.end = len(line) - 1
			value.val, _ = strconv.Atoi(prev)
			prev = ""
			vals = append(vals, value)
		}

		line_no++

		if err != nil {
			break
		}
	}

	sum := 0
	/*For Part - 1
	for _, val := range vals {
		start, end, line := val.start, val.end, val.line
		var pos_pairs []pair
		for i := line - 1; i <= line+1; i++ {
			if i == line {
				pos_pairs = append(pos_pairs, pair{line: i, pos: start - 1}, pair{line: i, pos: end + 1})
			} else {
				for j := start - 1; j <= end+1; j++ {
					pos_pairs = append(pos_pairs, pair{line: i, pos: j})
				}
			}
		}

		for _, pos := range pos_pairs {
			if _, check := sym[pos]; check {
				sum += val.val
			}
		}
	}*/

	for key, _ := range sym {
		var pos_pairs []pair
		for i := key.line - 1; i <= key.line+1; i++ {
			if i == key.line {
				pos_pairs = append(pos_pairs, pair{line: i, pos: key.pos - 1}, pair{line: i, pos: key.pos + 1})
			} else {
				for j := key.pos - 1; j <= key.pos+1; j++ {
					pos_pairs = append(pos_pairs, pair{line: i, pos: j})
				}
			}
		}
		gears, prod := 0, 1
		for _, val := range vals {
			for _, pos := range pos_pairs {
				if pos.pos >= val.start && pos.pos <= val.end && pos.line == val.line {
					gears++
					prod *= val.val
					break
				}
			}
		}
		if gears == 2 {
			sum += prod
		}

	}

	fmt.Println(sum)

}
