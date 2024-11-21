package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

const (
	userAgent      = "script:Search r/adventofcode:1.0 (by /u/arkteck)"
	accessTokenURL = "https://www.reddit.com/api/v1/access_token"
	searchURL      = "https://oauth.reddit.com/r/adventofcode/search.json"
)

func main() {
	getSolutionURLs()

}

func getSolutionURLs() {

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

	searchTerm := "flair:solution"
	resp, err := client.Get(searchURL + "?q=" + url.QueryEscape(searchTerm) + "&sort=new&restrict_sr=on&limit=1000")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch results: %v", resp.Status)
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

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range result.Data.Children {
		fmt.Printf("Title: %s\n", post.Data.Title)
		fmt.Printf("URL: %s\n", post.Data.URL)
	}
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
