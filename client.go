package bitbucket

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var ErrRateLimited = errors.New("bitbucket: rate limited")
var ErrNotFound = errors.New("bitbucket: record not found")
var ErrInternalServerError = errors.New("bitbucket: internal server error")

type Client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: httpClient,
	}
}

func (c *Client) GetURL(urlStr string) ([]byte, error) {
	resp, err := c.client.Get(urlStr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, ErrRateLimited
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}
	if resp.StatusCode == http.StatusInternalServerError {
		return nil, ErrInternalServerError
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got %s, wanted 200 OK", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) getJSON(urlStr string, v interface{}) error {
	contents, err := c.GetURL(urlStr)
	if err != nil {
		return err
	}
	return json.Unmarshal(contents, v)
}
