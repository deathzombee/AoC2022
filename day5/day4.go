package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve() (string, string) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var stackA [][]string
	var topValuesA string
	var topValuesB string

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		chars := strings.Split(line, "")
		//fmt.Println(chars)

		if len(chars) == 0 {
			break
		}

		for i, char := range chars {
			if char == "[" {
				pos := i / 4
				//fmt.Println("pos", pos)

				if len(stackA) <= pos {
					// extending stack
					stackA = append(stackA, make([][]string, pos-len(stackA)+1)...)
					fmt.Print("stackA", stackA)
				}

				stackA[pos] = append(stackA[pos], chars[i+1])
				fmt.Println("stackA", stackA)
			}
		}
	}
	stackB := make([][]string, len(stackA))
	copy(stackB, stackA)
	for scanner.Scan() {
		lineh := scanner.Text()

		var amt, src, tgt int
		fmt.Sscanf(lineh, "move %d from %d to %d", &amt, &src, &tgt, stackA, stackB)
		stackmover(amt, src, tgt, stackA, stackB)

	}

	for _, sub := range stackA {
		topValuesA += sub[0]
	}

	for _, sub := range stackB {
		topValuesB += sub[0]
	}

	return topValuesA, topValuesB
}
func stackmover(amt int, src int, tgt int, stackA [][]string, stackB [][]string) {

	for i := 0; i < amt; i++ {
		var crate string
		crate, stackA[src-1] = stackA[src-1][0], stackA[src-1][1:]
		stackA[tgt-1] = append([]string{crate}, stackA[tgt-1]...)
		fmt.Println("stackA", stackA)

		crate, stackB[src-1] = stackB[src-1][0], stackB[src-1][1:]
		if len(stackB[tgt-1]) == i {
			stackB[tgt-1] = append(stackB[tgt-1], crate)
		} else {
			stackB[tgt-1] = append(stackB[tgt-1][:i+1], stackB[tgt-1][i:]...)
			stackB[tgt-1][i] = crate
		}
	}
}
func main() {
	result_a, result_b := solve()
	fmt.Printf("A: %s\nB: %s\n", result_a, result_b)
}
