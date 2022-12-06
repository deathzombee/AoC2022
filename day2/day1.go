package main

import (
	"bufio"
	_ "day2/grab"
	"fmt"
	"log"
	_ "log"
	"os"
	"strings"
)

// A Y
// B X
// C Z
func rps(a string, b string) (string, int, int) {
	//points each type of character is worth
	var moveMap1 = map[string]int{
		"A": 0, //rock
		"B": 1, //paper
		"C": 2, //scissors
	}
	var moveMap2 = map[string]int{
		"X": 0, //rock or lose
		"Y": 1, //paper or draw
		"Z": 2, //scissors or win
	}
	// to figure out the appropriate move for a given outcome
	reversematrix := [3][3]int{
		{2, 0, 1},
		{0, 1, 2},
		{1, 2, 0},
	}
	matrix := [3][3]int{
		{0, -1, 1},
		{1, 0, -1},
		{-1, 1, 0},
	}
	pointmatrix := [3][3]int{
		{4, 8, 3},
		{1, 5, 9},
		{7, 2, 6},
	}
	p1 := moveMap1[a]
	p2 := moveMap2[b]
	//reimplwment this
	//calc := matrix[p1][p2]
	rb := reversematrix[p1][p2]
	reversecalc := matrix[p1][rb]
	if reversecalc == 0 {
		return "Tie", pointmatrix[p1][rb], pointmatrix[rb][p1]
	}
	if reversecalc == 1 {
		return "Player 1 wins", pointmatrix[p1][rb], pointmatrix[rb][p1]
	}
	if reversecalc == -1 {
		return "Player 2 wins", pointmatrix[p1][rb], pointmatrix[rb][p1]
	}
	return "error", 0, 0
}

func main() {
	//if _, err := os.Stat("input.txt"); os.IsNotExist(err) {
	//	log.Println("input.txt does not exist, running grab.go")
	//	grab.Grab("4")
	//
	//}
	//read input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	var a []string
	var b []string
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			a = append(a, strings.ToUpper(string(input[i])))
		} else {
			b = append(b, strings.ToUpper(string(input[i])))
		}

	}
	accu := 0
	for i := 0; i < len(a); i++ {
		_, t, c := rps(a[i], b[i])
		fmt.Println(a[i], b[i], t, c)
		//fmt.Println(t)
		accu += t
	}
	fmt.Println(accu)
}
