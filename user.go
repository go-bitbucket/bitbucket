package bitbucket

import (
	"fmt"
	"time"
)

type User struct {
	UUID        string `json:"uuid"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	Links       struct {
		Avatar Link `json:"avatar"`
		HTML   Link `json:"html"`
		Self   Link `json:"self"`
	}
}

type UserProfile struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	Website     string `json:"website"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Links       UserProfileLinks
	CreatedOn   time.Time `json:"created_on"`
}

type EmailsPage struct {
	Page    int `json:"page"`
	PageLen int `json:"pagelen"`
	Size    int `json:"size"`
	Values  []Email
}

type Email struct {
	Email       string `json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
	IsPrimary   bool   `json:"is_primary"`
	Type        string `json:"type"`
}

type UserProfileLinks struct {
	Self         Link
	Repositories Link
	HTML         Link
	Followers    Link
	Avatar       Link
	Following    Link
}

func UserURL(user string) string {
	return fmt.Sprintf(
		"https://api.bitbucket.org/2.0/users/%s",
		user,
	)
}

func TeamURL(team string) string {
	return fmt.Sprintf(
		"https://api.bitbucket.org/2.0/teams/%s",
		team,
	)
}

func (c *Client) GetUser(urlStr string) (*UserProfile, error) {
	var user UserProfile
	err := c.getJSON(urlStr, &user)
	return &user, err
}

func (c *Client) GetEmail() (*EmailsPage, error) {
	var email EmailsPage
	err := c.getJSON("https://api.bitbucket.org/2.0/user/emails", &email)
	return &email, err
}
