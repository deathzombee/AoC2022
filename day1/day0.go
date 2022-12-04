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
