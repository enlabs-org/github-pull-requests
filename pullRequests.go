package GithubPullRequests

import (
	"github.com/google/go-github/v58/github"
	"golang.org/x/net/context"
	"log"
)

func GetPullRequestNumbers(Config GithubConfiguration, State string) ([]int, error) {
	client := github.NewClient(nil)
	if Config.authToken != nil {
		client.WithAuthToken(*Config.authToken)
	} 
	
	pulls, _, err := client.PullRequests.List(context.Background(), Config.owner, Config.repo, nil)
	
	var prNumbers []int
	for _, pull := range pulls {
		if *pull.State == State {
			prNumbers = append(prNumbers, *pull.Number)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	
	return prNumbers, err
}
