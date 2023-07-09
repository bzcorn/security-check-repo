package checker

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func CheckRepos(repos []string, token string) {
	// Create a context
	ctx := context.Background()

	// Create a client with your token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Iterate over the list of repos
	for _, repo := range repos {
		splitted := splitRepoName(repo)

		commits, _, err := client.Repositories.ListCommits(ctx, splitted[0], splitted[1], &github.CommitsListOptions{})
		if err != nil {
			fmt.Printf("Error retrieving commits for repo: %v\n", err)
			continue
		}

		// Get the latest commit
		latestCommit := commits[0]
		if latestCommit.Commit.Committer.Date.After(time.Now().AddDate(-1, 0, 0)) {
			fmt.Printf("Repo %s has a commit in the last year\n", repo)
		} else {
			fmt.Printf("Repo %s hasn't had a commit in the last year\n", repo)
		}
	}
}

func splitRepoName(repo string) []string {
	splitted := strings.Split(repo, "/")
	if len(splitted) != 2 {
		panic("Invalid repo format, expected 'owner/repo'")
	}
	return splitted
}
