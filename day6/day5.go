package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//find start of packet signal

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	//buffer the input
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 1024*1024), 2)
	var ret []byte
	for scanner.Scan() {
		ret = append(ret, scanner.Bytes()...)
		var eh []byte
		for i := 0; i < len(ret)-4; i++ {
			eh = ret[i : i+14]
			//check if each character is unique
			if unique(eh) {
				fmt.Println(i + 14)
				fmt.Println(string(eh))

				break
			} else {

			}

		}
		fmt.Println(string(ret[18]))
		fmt.Println(string(ret[24]))

	}

}

// jpqm g marker g
// qmgbljsphdztnv j marker j
func unique(s []byte) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
