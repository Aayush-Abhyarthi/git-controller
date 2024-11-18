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

	// Authenticate using a personal access token (set in environment variable)
	var token string
	fmt.Scan(&token)
	//token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN environment variable is not set")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	// Create a new GitHub client
	client := github.NewClient(tc)

	// Fetch the organizations
	orgs, _, err := client.Organizations.List(ctx, "", nil)
	if err != nil {
		log.Fatalf("Error fetching organizations: %v", err)
	}

	// Print the organizations
	if len(orgs) == 0 {
		fmt.Println("No organizations found.")
		return
	}

	fmt.Println("Organizations your account has access to:")
	for _, org := range orgs {
		fmt.Printf("- %s (URL: %s)\n", *org.Login, *org.HTMLURL)
	}
}