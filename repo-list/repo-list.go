package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	// Use a personal access token (replace with your token)
	var token string
	fmt.Scan(&token)
	//token := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Example: List your repositories
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Fatalf("Error listing repos: %v", err)
	}

	for _, repo := range repos {
		log.Printf("Repo: %s\n", *repo.Name)
	}
}
