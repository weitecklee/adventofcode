package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	userAgent      = "script:Search r/adventofcode:1.0 (by /u/arkteck)"
	accessTokenURL = "https://www.reddit.com/api/v1/access_token"
	searchURL      = "https://oauth.reddit.com/r/adventofcode/search.json"
)

func main() {

	var year, day string
	fmt.Println("Enter year and day (e.g., \"2020 20\"):")
	fmt.Scanln(&year, &day)
	solutionURL := getSolutionURL(year, day)
	createReadMeFile(solutionURL)
	getInput(year, day)
	if err := copyTemplates(year, day); err != nil {
		log.Fatalf("Failed to copy templates: %v", err)
	}
}

func getInput(year, day string) {

	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")
	if sessionCookie == "" {
		fmt.Println("Error: AOC_SESSION_COOKIE environment variable not set")
		return
	}

	url := "https://adventofcode.com/" + year + "/day/" + day + "/input"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))

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
	filePath := filepath.Join(dir, "input.txt")

	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("Input saved to input.txt")
}

func getSolutionURL(year, day string) SolutionURL {

	log.SetFlags(log.Llongfile)

	clientID := os.Getenv("REDDIT_CLIENT_ID")
	if clientID == "" {
		log.Fatal("Error: REDDIT_CLIENT_ID environment variable not set")
	}
	clientSecret := os.Getenv("REDDIT_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("Error: REDDIT_CLIENT_SECRET environment variable not set")
	}

	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     accessTokenURL,
		Scopes:       []string{"read"},
	}

	client := config.Client(context.Background())
	client.Transport = &userAgentTransport{
		Transport: client.Transport,
		UserAgent: userAgent,
	}

	var result struct {
		Data struct {
			Children []struct {
				Data struct {
					Title string `json:"title"`
					URL   string `json:"url"`
				} `json:"data"`
			} `json:"children"`
		} `json:"data"`
	}

	searchTerm := fmt.Sprintf(`flair:solution title:"%s day %s`, year, day)
	resp, err := client.Get(searchURL + "?q=" + url.QueryEscape(searchTerm) + "&sort=new&restrict_sr=on&limit=1000")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch results: %v", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	solutionURL := SolutionURL{
		Title: result.Data.Children[0].Data.Title,
		URL:   result.Data.Children[0].Data.URL,
		Year:  year,
		Day:   day,
	}

	return solutionURL
}

func createReadMeFile(solutionURL SolutionURL) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	year, day, URL := solutionURL.Year, solutionURL.Day, solutionURL.URL
	dirPath := filepath.Join(wd, year, day)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatalf("Failed to create directories: %v", err)
	}
	filePath := filepath.Join(dirPath, "README.md")
	puzzleURL := fmt.Sprintf("https://adventofcode.com/%s/day/%s", year, day)
	body := fmt.Sprintf("### Advent of Code %s Day %s\n\n[Puzzle Page](%s)\n\n[Solutions Megathread](%s)\n", year, day, puzzleURL, URL)
	err = os.WriteFile(filePath, []byte(body), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

}

type SolutionURL struct {
	Title string
	Year  string
	Day   string
	URL   string
}

// userAgentTransport adds a custom User-Agent header to all HTTP requests
type userAgentTransport struct {
	Transport http.RoundTripper
	UserAgent string
}

func (u *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", u.UserAgent)
	return u.Transport.RoundTrip(req)
}

func copyTemplates(year, day string) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return err
	}
	srcFolder := filepath.Join(wd, "templates")
	dstFolder := filepath.Join(wd, year, day)

	return filepath.Walk(srcFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcFolder, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(dstFolder, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		} else {
			return copyFile(path, destPath)
		}
	})
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, sourceInfo.Mode())
}
