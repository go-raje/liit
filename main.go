package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println(Hello("Reu"))
}

func Hello(person string) string {
	return "Hello, " + person
}

func DownloadFile(filepath string, url string) error {
	// hit the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// copy response body for GET url onto created file
	_, err = io.Copy(file, resp.Body)
	return err

}
