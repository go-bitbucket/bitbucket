package bitbucket

import (
	"fmt"
	"time"
)

type CommitsPage struct {
	PageLen int      `json:"pagelen"`
	Values  []Commit `json:"values"`
	Next    string   `json:"next"`
}

type Commit struct {
	Hash       string      `json:"hash"`
	Repository CommitRepo  `json:"repository"`
	Links      CommitLinks `json:"links"`
	Author     Author
	Parents    []struct {
		Hash  string `json:"hash"`
		Type  string `json:"type"`
		Links struct {
			Self Link `json:"self"`
			HTML Link `json:"html"`
		} `json:"links"`
	} `json:"parents"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}

type CommitAuthor struct {
	UUID        string `json:"uuid"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Links       struct {
		Avatar Link `json:"avatar"`
		Self   Link `json:"self"`
	}
}

type Author struct {
	Raw  string `json:"raw"`
	User User   `json:"user"`
}

type CommitRepo struct {
	Links struct {
		Self   Link `json:"self"`
		HTML   Link `json:"html"`
		Avatar Link `json:"avatar"`
	} `json:"links"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	UUID     string `json:"uuid"`
}

type CommitLinks struct {
	Self     Link `json:"self"`
	Comments Link `json:"comments"`
	Patch    Link `json:"patch"`
	HTML     Link `json:"html"`
	Diff     Link `json:"diff"`
	Approve  Link `json:"approve"`
}

func CommitsURL(user, repo string) string {
	return fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/commits",
		user, repo,
	)
}

func CommitURL(user, repo, revision string) string {
	return fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s/commit/%s",
		user, repo, revision,
	)
}

func (c *Client) GetCommit(urlStr string) (*Commit, error) {
	var commit Commit
	err := c.getJSON(urlStr, &commit)
	return &commit, err
}

func (c *Client) ListCommits(urlStr string) (*CommitsPage, error) {
	var commitsPage CommitsPage
	err := c.getJSON(urlStr, &commitsPage)
	return &commitsPage, err
}
