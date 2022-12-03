package grab

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)


func Grab() {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go


	req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/1/input", nil)
	if err != nil {
		// handle err
	}
	//open cookies.txt and read it into a string
	content, err := os.ReadFile("cookies.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:107.0) Gecko/20100101 Firefox/107.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Referer", "https://adventofcode.com/2022/day/1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "session=" + string(content))
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Te", "trailers")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	//new,_ := io.ReadAll(resp.Body)
	gzreader,_ := gzip.NewReader(resp.Body)
	output,_ := io.ReadAll(gzreader)
	os.WriteFile("../input.txt",output,0644)


	fmt.Println(string(output))



}