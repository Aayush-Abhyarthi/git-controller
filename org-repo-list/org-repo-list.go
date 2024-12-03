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
	// token := os.Getenv("GITHUB_TOKEN")
	// if token == "" {
	// 	log.Fatal("GITHUB_TOKEN environment variable is not set")
	// }

	var token string
	fmt.Scan(&token)

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	// Create a new GitHub client
	client := github.NewClient(tc)

	// Replace with your organization's name
	var orgName string // Update this
	fmt.Scan(&orgName)

	// List repositories for the organization
	opt := &github.RepositoryListByOrgOptions{
		Type: "all", // Options: "all", "public", "private", "forks", "sources", "member"
		ListOptions: github.ListOptions{
			PerPage: 50, // Pagination size
		},
	}

	fmt.Printf("Repositories in organization: %s\n", orgName)

	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, orgName, opt)
		if err != nil {
			log.Fatalf("Error listing repositories: %v", err)
		}

		// Print repository names and details
		for _, repo := range repos {
			if repo.Name != nil && repo.HTMLURL != nil {
				fmt.Printf("- %s (URL: %s)\n", *repo.Name, *repo.HTMLURL)
			} else {
				fmt.Println("- Repository data missing or incomplete")
			}
		}

		// Check if there are more pages
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}
