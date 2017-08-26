package gogs

import (
	gogsclient "github.com/gogits/go-gogs-client"
)

type Config struct {
	URL   string
	Token string
}

func (c *Config) loadAndValidate() (*gogsclient.Client, error) {
	client := gogsclient.NewClient(c.URL, c.Token)

	_, err := client.ListMyRepos()
	if err != nil {
		return nil, err
	}

	return client, nil
}
