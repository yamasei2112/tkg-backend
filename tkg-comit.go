package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Login string `json:"login"`
}

type Repository struct {
	Name string `json:"name"`
}

type Contributor struct {
	Login string `json:"login"`
}

type Event struct {
	Type string `json:"type"`
}

func main() {
	username := "yamasei2112"
	user, err := getUserInfo(username)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Username: %s\n", user.Login)

	repos, err := getUserRepos(username)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	for _, repo := range repos {
		fmt.Printf("Repository: %s\n", repo.Name)

		contributors, err := getRepoContributors(username, repo.Name)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		for _, contributor := range contributors {
			fmt.Printf("  Contributor: %s\n", contributor.Login)
		}
	}
	events, err := getUserEvents(username)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Events for user %s:\n", username)
	for _, event := range events {
		fmt.Println(event.Type)
	}
}

func getUserInfo(username string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func getUserRepos(username string) ([]Repository, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func getRepoContributors(owner, repo string) ([]Contributor, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var contributors []Contributor
	if err := json.NewDecoder(resp.Body).Decode(&contributors); err != nil {
		return nil, err
	}

	return contributors, nil
}

func getUserEvents(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}
