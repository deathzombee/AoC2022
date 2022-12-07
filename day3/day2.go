package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lowermap := make(map[string]int)
	uppermap := make(map[string]int)
	fullmap := make(map[string]int)
	for i := 1; i <= 26; i++ {
		lowermap[string(i+96)] = i
		uppermap[string(i+64)] = i + 26
	}
	//combine the two maps
	for k, v := range uppermap {
		fullmap[k] = v
	}
	for k, v := range lowermap {
		fullmap[k] = v
	}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	//var match []int
	var inter []int
	//scan 3 lines at a time
	for scanner.Scan() {
		line1 := scanner.Text()

		scanner.Scan()
		line2 := scanner.Text()

		scanner.Scan()
		line3 := scanner.Text()

		inter = append(inter, mch(line1, line2, line3, fullmap))

	}
	//reimplement for solution 1
	//for scanner.Scan() {
	//
	//	rucksac := scanner.Text()
	//	rucksize := len(rucksac)
	//	firsthalf := rucksac[:rucksize/2]
	//	secondhalf := rucksac[rucksize/2:]
	//	var first = 0
	//
	//	for k := 0; k < len(firsthalf); k++ {
	//
	//		for j := 0; j < len(secondhalf); j++ {
	//
	//			//only match the same letter once
	//			if firsthalf[k] == secondhalf[j] && first == 0 {
	//				first += 1
	//				match = append(match, fullmap[string(firsthalf[k])])
	//				break
	//
	//			}
	//		}
	//	}
	//}
	//var sum int
	//for i := 0; i < len(match); i++ {
	//	sum += match[i]
	//}
	//fmt.Println(sum)
	var sum int
	for i := 0; i < len(inter); i++ {
		sum += inter[i]
	}
	fmt.Println(sum)
}
func mch(firsthalf string, secondhalf string, thirdhalf string, fullmap map[string]int) int {
	var first = 0
	var mach int
	for k := 0; k < len(firsthalf); k++ {
		// find match in all three lines
		for j := 0; j < len(secondhalf); j++ {
			for l := 0; l < len(thirdhalf); l++ {
				//only match the same letter once
				if firsthalf[k] == secondhalf[j] && firsthalf[k] == thirdhalf[l] && first == 0 {
					first += 1
					//mach = append(mach, fullmap[string(firsthalf[k])])
					mach = fullmap[string(firsthalf[k])]
					break

				}
			}
		}

	}
	return mach
}
