package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fp := os.Getenv("FILE_PATH")
	out, err := os.Create(fp)
	if err != nil {
		fmt.Println("Error occurred while creating a file.", err)
	}
	defer out.Close()

	url := os.Getenv("URL")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred while creating a file.", err)
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error occurred while saving the contents.", err)
	}

	fmt.Println("File downloaded successfully")
	os.Exit(0)
}
