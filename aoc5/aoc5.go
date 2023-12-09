package aoc5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type agromap struct {
	dest, src, length int
}

type agro_flow struct {
	from, to string
}

func destination(from int, vals []agromap) int {
	for _, val := range vals {
		if from >= val.src && from < val.src+val.length {
			return val.dest + from - val.src
		}
	}
	return from
}

func Aoc5() {
	//logfile, _ := os.Create("log.txt")
	//log.SetOutput(logfile)
	time_start := time.Now()
	input, _ := os.Open("input.txt")
	reader := bufio.NewReader(input)
	var seeds []int
	agro_info := map[agro_flow][]agromap{
		agro_flow{from: "seed", to: "soil"}:            []agromap{},
		agro_flow{from: "soil", to: "fertilizer"}:      []agromap{},
		agro_flow{from: "fertilizer", to: "water"}:     []agromap{},
		agro_flow{from: "water", to: "light"}:          []agromap{},
		agro_flow{from: "light", to: "temperature"}:    []agromap{},
		agro_flow{from: "temperature", to: "humidity"}: []agromap{},
		agro_flow{from: "humidity", to: "location"}:    []agromap{},
	}
	var current agro_flow
	for {
		line, err := reader.ReadString('\n')
		line = line[:len(line)-1]

		reg1 := regexp.MustCompile("(?i)[A-Za-z]+-to-[A-Za-z]+")
		reg2 := regexp.MustCompile("(?i)[0-9]+ [0-9]+ [0-9]+")

		if strings.Contains(line, "seeds:") {
			seed_line := line[strings.Index(line, ":")+1:]
			for _, val := range strings.Split(seed_line, " ") {
				if val != "" {
					num, _ := strconv.Atoi(val)
					seeds = append(seeds, num)
				}
			}
		} else if reg1.MatchString(line) {
			match := reg1.FindString(line)
			current = agro_flow{
				from: strings.Split(match, "-to-")[0],
				to:   strings.Split(match, "-to-")[1],
			}

		} else if reg2.MatchString(line) {
			var nums []int
			for _, val := range strings.Split(line, " ") {
				if val != "" {
					num, _ := strconv.Atoi(val)
					nums = append(nums, num)
				}
			}
			if len(agro_info[current]) == 0 {
				agro_info[current] = append(agro_info[current],
					agromap{
						dest:   nums[0],
						src:    nums[1],
						length: nums[2],
					},
				)
			} else {
				if nums[2]+nums[1] > agro_info[current][0].src+agro_info[current][0].length {
					var agro_temp []agromap = []agromap{agromap{
						dest:   nums[0],
						src:    nums[1],
						length: nums[2],
					}}
					agro_info[current] = append(agro_temp, agro_info[current]...)
				} else {
					agro_info[current] = append(agro_info[current],
						agromap{
							dest:   nums[0],
							src:    nums[1],
							length: nums[2],
						},
					)
				}
			}

		} else {
			current = agro_flow{from: "", to: ""}
		}

		if err != nil {
			break
		}
	}

	/*for key, value := range agro_info {
		fmt.Println(key)
		for _, val := range value {
			fmt.Println(val)
		}
		fmt.Println("next")
	}*/

	var min_dest int = -1
	map_keys := map[int]int{}

	for idx := 0; idx < len(seeds); idx += 2 {
		_, ok := map_keys[seeds[idx]]
		if ok {
			continue
		}
		for seed := seeds[idx]; seed < seeds[idx]+seeds[idx+1] && !ok; seed++ {
			dest_value := destination(seed, agro_info[agro_flow{from: "seed", to: "soil"}])
			dest_value = destination(dest_value, agro_info[agro_flow{from: "soil", to: "fertilizer"}])
			dest_value = destination(dest_value, agro_info[agro_flow{from: "fertilizer", to: "water"}])
			dest_value = destination(dest_value, agro_info[agro_flow{from: "water", to: "light"}])
			dest_value = destination(dest_value, agro_info[agro_flow{from: "light", to: "temperature"}])
			dest_value = destination(dest_value, agro_info[agro_flow{from: "temperature", to: "humidity"}])
			dest_value = destination(dest_value, agro_info[agro_flow{from: "humidity", to: "location"}])
			//log.Println(seed, dest_value, idx)
			map_keys[seeds[idx]] = dest_value
			if min_dest == -1 || dest_value < min_dest {
				min_dest = dest_value
			}
			//7326078
		}
	}
	//log.Println(min_dest)
	fmt.Println(min_dest)
	fmt.Println("time:", time.Since(time_start))

}
