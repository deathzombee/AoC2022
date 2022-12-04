package main

import (
	"bufio"
	"day4/grab"
	"github.com/JeremyOne/subsets"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// if input.txt does not exist run grab.go
	if _, err := os.Stat("input.txt"); os.IsNotExist(err) {
		log.Println("input.txt does not exist, running grab.go")
		grab.Grab("4")

	}
	//read input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()
	bufio.NewReader(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	glbAccu := 0
	accu2 := []int{}
	for i := 0; i < len(ret); i++ {
		if sat(ret[i], i, 0, 1) == 1 {
			glbAccu++
			accu2 = append(accu2, i)
		}
	}
	//for numbers not in accu2 rerun scan backwards
	for i := 0; i < len(ret); i++ {
		if !contains(accu2, i) {
			if sat(ret[i], i, 1, 0) == 1 {
				glbAccu++
				accu2 = append(accu2, i)
			}
		}
	}
	println(glbAccu)

}

func contains(accu2 []int, i int) bool {
	for _, v := range accu2 {
		if v == i {
			return true
		}
	}
	return false

}

func sat(s string, b int, c int, d int) int {
	accu := 0
	rp := strings.Split(s, ",")
	rp1 := strings.Split(rp[0], "-")
	rp2 := strings.Split(rp[1], "-")
	rinp1 := make([]int, len(rp1))
	range1list := make([]string, 0)
	for i := 0; i < len(rp1); i++ {
		rinp1[i], _ = strconv.Atoi(rp1[i])
	}
	for i := rinp1[0]; i <= rinp1[1]; i++ {
		range1list = append(range1list, string(rune(i)))
	}
	rinp2 := make([]int, len(rp2))
	range2list := make([]string, 0)
	for i := 0; i < len(rp1); i++ {
		rinp2[i], _ = strconv.Atoi(rp2[i])
	}
	for i := rinp2[0]; i <= rinp2[1]; i++ {
		range2list = append(range2list, string(rune(i)))
	}
	var superrange [2][]string
	superrange[0] = range1list
	superrange[1] = range2list
	found := subsets.ArrayIsSubset_BruteForce(superrange[c], superrange[d], true)
	if found {
		accu = 1
	} else {
		accu = 0
	}
	return accu
}
