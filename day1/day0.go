package main

import (
	"bufio"
	"day1/grab"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// if input.txt does not exist run grab.go
	if _, err := os.Stat("input.txt"); os.IsNotExist(err) {
		log.Println("input.txt does not exist, running grab.go")
		grab.Grab("1")

	}

	ret := tb()
	fmt.Println(maximum(ret))
	fmt.Println(top3sum(ret))

}
func top3sum(a []int) int {
	m1 := maximum(a)
	a = remove(a, m1)
	m2 := maximum(a)
	a = remove(a, m2)
	m3 := maximum(a)
	return m1 + m2 + m3
}

func remove(a []int, m1 interface{}) []int {
	for i, v := range a {
		if v == m1 {
			a = append(a[:i], a[i+1:]...)
		}
	}
	return a

}
func maximum(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func tb() []int {
	os.Open("input.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 2)
	var ret []byte
	accum := 0
	var arr []int
	for scanner.Scan() {
		ret = append(ret, scanner.Bytes()...)
		if scanner.Text() == "" {
			//add to array
			arr = append(arr, accum)
			accum = 0
		} else {
			t, _ := strconv.Atoi(scanner.Text())
			accum += t
		}

	}
	return arr
}
