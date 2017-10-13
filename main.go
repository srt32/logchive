package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("PAPERTRAIL_API_KEY")
	if token == "" {
		fmt.Println("PAPERTRAIL_API_KEY is required")
		return
	}

	archiveURL := "https://papertrailapp.com/api/v1/archives/%v/download"

	url := fmt.Sprintf(archiveURL, "2017-09-24-14")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Papertrail-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("http error: %s", resp.Status)
	}
	defer resp.Body.Close()

	out, err := os.Create("new_file.tsv")
	if err != nil {
		fmt.Println("creating file error: %v", err)
	}

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("body error: %v", err)
	}

	fmt.Println("bytes downloaded: %v", n)
}
