package aoc7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func type_hand(str string) int {
	chars := make(map[rune]int)
	for _, char := range str {
		_, ok := chars[char]
		if ok {
			chars[char]++
		} else {
			chars[char] = 1
		}
	}
	keys := make([]rune, 0, len(chars))
	for key, _ := range chars {
		keys = append(keys, key)
	}
	numkeys := len(chars)

	if numkeys == 1 {
		return 6
	} else if numkeys == 2 {
		if chars[keys[0]]/chars[keys[1]] == 4 || chars[keys[1]]/chars[keys[0]] == 4 {
			if chars['J'] > 0 {
				return 6
			}
			return 5
		} else {
			if chars['J'] > 0 {
				return 6
			}
			return 4
		}
	} else if numkeys == 3 {
		if chars[keys[0]] == 2 || chars[keys[1]] == 2 || chars[keys[2]] == 2 {
			if chars['J'] == 1 {
				return 4
			}
			if chars['J'] == 2 {
				return 5
			}
			return 2
		} else {
			if chars['J'] == 1 || chars['J'] == 3 {
				return 5
			}
			return 3
		}
	} else if numkeys == 4 {
		if chars['J'] > 0 {
			return 3
		}
		return 1
	} else {
		if chars['J'] > 0 {
			return 1
		}
		return 0
	}

}

func compare_hands(hand1 string, hand2 string) bool {
	mapping := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		//"J": 11,
		"T": 11,
		"9": 10,
		"8": 9,
		"7": 8,
		"6": 7,
		"5": 6,
		"4": 5,
		"3": 4,
		"2": 3,
		"J": 2,
	}
	//fmt.Println(hand1, hand2)
	for i := range hand1 {
		//fmt.Println(hand1[i:i+1], hand2[i:i+1])
		if mapping[hand1[i:i+1]] > mapping[hand2[i:i+1]] {
			//fmt.Println("true")
			return true
		} else if mapping[hand1[i:i+1]] < mapping[hand2[i:i+1]] {
			//fmt.Println("false")
			return false
		}
	}
	return true
}

func sort_lists(hand int, str string, sorted_list map[int][]string) {
	_, ok := sorted_list[hand]
	i, j := 0, len(sorted_list[hand])-1
	if ok {

		for i <= j {
			mid := (i + j) / 2
			if compare_hands(str, sorted_list[hand][mid]) {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}

	} else {
		sorted_list[hand] = append(sorted_list[hand], str)
		return
	}
	helper1 := make([]string, len(sorted_list[hand]))
	_ = copy(helper1, sorted_list[hand])
	helper := append(sorted_list[hand][0:i], str)
	sorted_list[hand] = append(helper, helper1[i:]...)
}

func Aoc7() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	sorted_list := make(map[int][]string)
	hand_bid := make(map[string]int)
	for {
		line, err := reader.ReadString('\n')

		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		vals := strings.Split(line, " ")
		hand := type_hand(vals[0])
		bid, _ := strconv.Atoi(vals[1])
		hand_bid[vals[0]] = bid

		sort_lists(hand, vals[0], sorted_list)

		if err != nil {
			break
		}
	}

	count, sum := 1, 0
	for i := 0; i <= 6; i++ {
		for j := len(sorted_list[i]) - 1; j >= 0; j-- {
			fmt.Println(sorted_list[i][j], i)
			sum += hand_bid[sorted_list[i][j]] * count
			count++
		}
	}

	fmt.Println(sum)
}
