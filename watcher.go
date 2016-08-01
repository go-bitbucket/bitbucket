package bitbucket

import (
	"fmt"
)

type WatchersPage struct {
	Page    int
	PageLen int
	Size    int
	Next    string
	Values  []Watcher
}

type Watcher struct {
	UUID        string
	Username    string
	DisplayName string
	Type        string
	Links       WatcherLinks
}

type WatcherLinks struct {
	Self   Link
	HTML   Link
	Avatar Link
}

func WatchersURL(user, repo string) string {
	return fmt.Sprintf(
		" https://api.bitbucket.org/2.0/repositories/%s/%s/watchers",
		user, repo,
	)
}

func (c *Client) ListWatchers(urlStr string) (*WatchersPage, error) {
	var watchersPage WatchersPage
	err := c.getJSON(urlStr, &watchersPage)
	return &watchersPage, err
}
