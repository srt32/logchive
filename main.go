package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("PAPERTRAIL_API_KEY")
	if token == "" {
		fmt.Printf("PAPERTRAIL_API_KEY is required")
		return
	}

	archiveURL := "https://papertrailapp.com/api/v1/archives/%v/download"

	url := fmt.Sprintf(archiveURL, "2017-09-24-14")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Papertrail-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http error: %s", resp.Status)
	}
	defer resp.Body.Close()

	out, err := os.Create("new_file.tsv")
	if err != nil {
		fmt.Printf("creating file error: %v", err)
	}

	unzippedArchive, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Printf("ungzipping error: %v", err)
	}
	defer unzippedArchive.Close()

	n, err := io.Copy(out, unzippedArchive)
	if err != nil {
		fmt.Printf("copying to file error: %v", err)
	}

	fmt.Printf("bytes downloaded: %v", n)
}
