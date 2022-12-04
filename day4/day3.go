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
	colli := 0
	glbAccu := 0
	accu2 := []int{}
	//collision is doesnt require backwards search so we can piggyback on the first loop
	// this saves a lot of calls to terser(), but this is a hack job so whatever
	// should probably make a data structure to hold all the returns from terser in one loop
	// and reuse it in the second loop
	for i := 0; i < len(ret); i++ {
		in := terser(ret[i])
		if collision(in[0], in[1]) {

			colli++

		}
		if sat(in, i, 0, 1) == 1 {
			glbAccu++
			accu2 = append(accu2, i)
		}
	}
	//for numbers not in accu2 rerun scan backwards
	for i := 0; i < len(ret); i++ {
		if !contains(accu2, i) {
			in := terser(ret[i])
			if sat(in, i, 1, 0) == 1 {
				glbAccu++
				accu2 = append(accu2, i)
			}
		}
	}
	println(glbAccu)
	println(colli)
}

func collision(s []string, e []string) bool {
	for _, v := range s {
		for _, v2 := range e {
			if v == v2 {
				return true
			}

		}
	}
	return false
}

func contains(accu2 []int, i int) bool {
	for _, v := range accu2 {
		if v == i {
			return true
		}
	}
	return false

}

func sat(s [2][]string, b int, c int, d int) int {
	accu := 0

	found := subsets.ArrayIsSubset_BruteForce(s[c], s[d], true)
	if found {
		accu = 1
	} else {
		accu = 0
	}
	return accu
}
func terser(s string) [2][]string {
	rp := strings.Split(s, ",")
	rp1 := strings.Split(rp[0], "-")
	rp2 := strings.Split(rp[1], "-")
	rinp1 := make([]int, len(rp1))
	rinp2 := make([]int, len(rp2))
	var superrange [2][]string
	for i := 0; i < len(rp1); i++ {
		rinp1[i], _ = strconv.Atoi(rp1[i])
	}
	for i := rinp1[0]; i <= rinp1[1]; i++ {
		superrange[0] = append(superrange[0], string(rune(i)))
	}
	for i := 0; i < len(rp1); i++ {
		rinp2[i], _ = strconv.Atoi(rp2[i])
	}
	for i := rinp2[0]; i <= rinp2[1]; i++ {
		superrange[1] = append(superrange[1], string(rune(i)))
	}

	return superrange
}
