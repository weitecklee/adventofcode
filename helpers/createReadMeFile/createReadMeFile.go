package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	userAgent      = "script:Search r/adventofcode:1.0 (by /u/arkteck)"
	accessTokenURL = "https://www.reddit.com/api/v1/access_token"
	searchURL      = "https://oauth.reddit.com/r/adventofcode/search.json"
)

func main() {
	solutionURLs := getSolutionURLs()
	for _, solutionURL := range solutionURLs {
		fmt.Printf("%s %s %s\n", solutionURL.Year, solutionURL.Day, solutionURL.URL)
	}
	createReadMeFile(solutionURLs)
	createReadMeFile2015()
}

func createReadMeFile2015() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	for day := 1; day < 26; day++ {
		dirPath := filepath.Join(wd, "2015", fmt.Sprintf("%d", day))
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			log.Fatalf("Failed to create directories: %v", err)
		}
		filePath := filepath.Join(dirPath, "README.md")
		puzzleURL := fmt.Sprintf("https://adventofcode.com/2015/day/%d", day)
		body := fmt.Sprintf("### Advent of Code 2015 Day %d\n\n[Puzzle Page](%s)\n\n[Solutions Megathread]()\n", day, puzzleURL)
		err = os.WriteFile(filePath, []byte(body), 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}

	}
}

func createReadMeFile(solutionURLs []SolutionURL) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	for _, solutionURL := range solutionURLs {
		year, day, URL := solutionURL.Year, solutionURL.Day, solutionURL.URL
		dirPath := filepath.Join(wd, year, day)
		err := os.MkdirAll(dirPath, 0755)
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
}

func getSolutionURLs() []SolutionURL {

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

	solutionURLs := []SolutionURL{}

	for year := 2016; year < 2024; year++ {

		searchTerm := fmt.Sprintf("flair:solution title:%d", year)
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

		for _, post := range result.Data.Children {
			reg := regexp.MustCompile(`\d+`)
			nums := reg.FindAllString(post.Data.Title, -1)
			solutionURLs = append(solutionURLs, SolutionURL{
				Title: post.Data.Title,
				URL:   post.Data.URL,
				Year:  nums[0],
				Day:   strings.TrimLeft(nums[1], "0"),
			})
		}
	}

	return solutionURLs
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
