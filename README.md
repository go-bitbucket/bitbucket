# go-bitbucket #

go-bitbucket is a Go client library for accessing the Bitbucket API.

### Authentication ###

Create a new Bitbucket client:

```go
import (
	"golang.org/x/oauth2"
	bb "golang.org/x/oauth2/bitbucket"
	"gopkg.in/bitbucket.v1"
)

config := &oauth2.Config{
	ClientID:     "client_id",
	ClientSecret: "client_secret",
	RedirectURL:  "http://redirect_url/callback",
	Endpoint:     bb.Endpoint,
}

client = conf.Client(nil)
bitbucketClient = bitbucket.NewClient(client)
```

### Example ###

```go
repos, err := bitbucketClient.ListRepos(bitbucket.ReposURL("username"))
if err != nil {
	panic err
}

repo, err := bitbucketClient.GetRepo(bitbucket.RepoURL("username", "reponame"))
if err != nil {
	panic err
}
```