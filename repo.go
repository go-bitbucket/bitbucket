package bitbucket

import (
	"fmt"
	"time"
)

const UserReposURL = "https://api.bitbucket.org/2.0/repositories/?role=member"

type ReposPage struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Page     int    `json:"page"`
	PageLen  int    `json:"pagelen"`
	Size     int    `json:"size"`
	Values   []Repo `json:"values"`
}

type Repo struct {
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Owner       User        `json:"owner"`
	IsPrivate   bool        `json:"is_private"`
	Description string      `json:"description"`
	ForkPolicy  string      `json:"fork_policy"`
	Website     string      `json:"website"`
	HasIssues   bool        `json:"has_issues"`
	HasWiki     bool        `json:"has_wiki"`
	Language    string      `json:"language"`
	Links       RepoLinks   `json:"links"`
	SCM         string      `json:"scm"`
	Size        int         `json:"size"`
	Type        string      `json:"type"`
	UUID        string      `json:"uuid"`
	Commits     CommitsPage `json:"commits"`
	CreatedOn   time.Time   `json:"created_on"`
	UpdatedOn   time.Time   `json:"updated_on"`
}

type RepoLinks struct {
	Avatar Link `json:"avatar"`
	Clone  []struct {
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"clone"`
	Commits      Link `json:"commits"`
	Downloads    Link `json:"downloads"`
	Forks        Link `json:"forks"`
	Hooks        Link `json:"hooks"`
	HTML         Link `json:"html"`
	PullRequests Link `json:"pullrequests"`
	Self         Link `json:"self"`
	Watchers     Link `json:"watchers"`
}

type Link struct {
	Href string `json:"href"`
}

func ReposURL(user string) string {
	return fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s",
		user,
	)
}

func RepoURL(user, repo string) string {
	return fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s",
		user, repo,
	)
}

func (c *Client) ListRepos(urlStr string) (*ReposPage, error) {
	var reposPage ReposPage
	err := c.getJSON(urlStr, &reposPage)
	return &reposPage, err
}

func (c *Client) GetRepo(urlStr string) (*Repo, error) {
	var repo Repo
	err := c.getJSON(urlStr, &repo)
	return &repo, err
}
