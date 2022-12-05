package main

import (
	"bufio"
	"day4/grab"
	"fmt"
	"golang.org/x/tools/container/intsets"
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
	collision := 0
	for i := 0; i < len(ret); i++ {
		 res := terser2(ret[i])
		 if res[0][0]{
			 glbAccu++
		 }
		 if res[1][0] {
			 collision++
		 }
	}
	fmt.Println(glbAccu, collision)
}

func terser2(s string)[2][]bool {
	rp := strings.Split(s, ",")
	rp1 := strings.Split(rp[0], "-")
	rp2 := strings.Split(rp[1], "-")
	rinp1 := make([]int, len(rp1))
	rinp2 := make([]int, len(rp2))
	var set1 intsets.Sparse
	var set2 intsets.Sparse

	for i := 0; i < len(rp1); i++{
		rinp1[i], _ = strconv.Atoi(rp1[i])
	}
	for i := rinp1[0]; i <= rinp1[1]; i++{
		set1.Insert(i)

	}
	for i := 0; i < len(rp1); i++{
		rinp2[i], _ = strconv.Atoi(rp2[i])
	}
	for i := rinp2[0]; i <= rinp2[1]; i++{

		set2.Insert(i)
	}

	var isSubset bool
	var intersect bool
	returnval := [2][]bool{}
	isSubset1 := set1.SubsetOf(&set2)
	isSubset2 := set2.SubsetOf(&set1)
	isIntersect := set1.Intersects(&set2)
	if isSubset1 || isSubset2{

		isSubset = true
	}
	if isIntersect{
		intersect = true
	}

	returnval[0] = append(returnval[0], isSubset)
	returnval[1] = append(returnval[1], intersect)
	return returnval
}


