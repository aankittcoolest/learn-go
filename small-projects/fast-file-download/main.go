package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var download_url string

	fmt.Print("Please enter a URL:")
	fmt.Scanf("%s", &download_url)

	// Get header from the URL
	log.Println("URL:", download_url)
	file_size, err := getSizeAndCheckRangeSupport(download_url)
	handleError(err)
	log.Printf("File size: %d bytes\n", file_size)

}

func handleError(err error) {
	if err != nil {
		log.Println("err:", err)
		os.Exit(1)
	}
}

func getSizeAndCheckRangeSupport(url string) (size int64, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	log.Printf("Response header: %v\n", res.Header)
	fmt.Println(res.Header)
	header := res.Header
	accept_ranges, supported := header["Accept-Ranges"]

	if !supported {
		return 0, errors.New("Doesn't support header `Accept-Ranges`.")
	} else if supported && accept_ranges[0] != "bytes" {
		return 0, errors.New("Support `Accept-Ranges` but not `bytes`.")
	}

	size, err = strconv.ParseInt(header["Content-Length"][0], 10, 64)
	return
}
