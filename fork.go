package bitbucket

import (
	"fmt"
	"time"
)

type ForksPage struct {
	Page    int
	PageLen int
	Size    int
	Next    string
	Values  []Fork
}

type Fork struct {
	UUID        string
	FullName    string
	Description string
	ForkPolicy  string
	Language    string
	Scm         string
	Size        int
	Type        string
	HasIssues   bool
	HasWiki     bool
	IsPrivate   bool
	Parent      ParentFork
	CreatedOn   time.Time
	UpdatedOn   time.Time
}

type ParentFork struct {
	UUID     string
	FullName string
	Name     string
	Type     string
}

func ForksURL(user, repo string) string {
	return fmt.Sprintf(
		"https://api.bitbucket.org/2.0/repositories/%s/%s/forks",
		user, repo,
	)
}

func (c *Client) ListForks(urlStr string) (*ForksPage, error) {
	var forksPage ForksPage
	err := c.getJSON(urlStr, &forksPage)
	return &forksPage, err
}
