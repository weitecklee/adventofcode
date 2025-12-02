package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	userAgentAOC = "github.com/weitecklee/adventofcode"
)

func main() {
	var year, day string
	fmt.Println("Enter year and day (e.g., \"2020 20\"):")
	fmt.Scanln(&year, &day)
	getInput(year, day)
}

func getInput(year, day string) {

	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")
	if sessionCookie == "" {
		fmt.Println("Error: AOC_SESSION_COOKIE environment variable not set")
		fmt.Println("Enter cookie value: ")
		fmt.Scanln(&sessionCookie)
		if sessionCookie == "" {
			return
		}
		os.Setenv("AOC_SESSION_COOKIE", sessionCookie)
	}

	url := "https://adventofcode.com/" + year + "/day/" + day + "/input"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))
	req.Header.Set("User-Agent", userAgentAOC)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	dir := filepath.Join(wd, year, day)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	filePath := filepath.Join(dir, "input.txt")

	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("Input saved to input.txt")
}
