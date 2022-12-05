package main

import (
	"bufio"
	_ "day2/grab"
	"fmt"
	_ "log"
	"os"
	"strings"
)

//A Y
//B X
//C Z
func rps(a string,b string) (string, int, int) {
	//points each type of character is worth
	var moveMap1 = map[string]int{
		"A": 0, //rock
		"B": 1, //paper
		"C": 2, //scissors
	}
	var moveMap2 = map[string]int{
		"X": 0, //rock
		"Y": 1, //paper
		"Z": 2, //scissors
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
	calc := matrix[p1][p2]
	if calc == 0 {
		return "Tie" , pointmatrix[p1][p2],pointmatrix[p2][p1]
	}
	if calc == 1 {
		return "Player 1 wins", pointmatrix[p1][p2],pointmatrix[p2][p1]
	}
	if calc == -1 {
		return "Player 2 wins" ,pointmatrix[p1][p2],pointmatrix[p2][p1]
	}
	return "error", 0,0
}
//rock paper scissors logic
//func rps(a, b string) string {
//	if points[a] == points[b] {
//		return "tie"
//	}
//	if points[a]-points[b] == -2 {
//		return a
//	} else if points[a]-points[b] == 2 {
//		return b
//	} else if points[a] > points[b] {
//		return a
//	} else {
//		return b
//	}
//}
func main () {
	//if _, err := os.Stat("input.txt"); os.IsNotExist(err) {
	//	log.Println("input.txt does not exist, running grab.go")
	//	grab.Grab("4")
	//
	//}
	////read input.txt
	//file, err := os.Open("input.txt")
	//if err != nil {
	//	log.Fatal(err)
	//
	//}
	//defer file.Close()
	//bufio.NewReader(file)
	//scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines)
	//var ret []string
	//for scanner.Scan() {
	//	ret = append(ret, scanner.Text())
	//}
	////split input.txt into 2 slices
	//var a []string
	//var b []string
	//for i := 0; i < len(ret); i++ {
	//	if i % 2 == 0 {
	//		a = append(a, ret[i])
	//	} else {
	//		b = append(b, ret[i])
	//	}
	//}
	//c
	//prompt user for input
	fmt.Println("Enter a string of characters")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
//split input into 2 slices
	var a []string
	var b []string
	for i := 0; i < len(input); i++ {
		if i % 2 == 0 {
			a = append(a, strings.ToUpper(string(input[i])))
		} else {
			b = append(b, strings.ToUpper(string(input[i])))
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(rps(a[i], b[i]))
	}
}