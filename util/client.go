package util

import (
	"fmt"
	rocket "github.com/krakowski/rocket/api"
	"os"
)

const (
	envHost = "ROCKET_HOST"
)

func NewRocketClient() (*rocket.Client, error) {
	serverUrl, present := os.LookupEnv(envHost)
	if !present {
		return nil, fmt.Errorf("ROCKET_HOST environment variable missing")
	}

	credentials := GetCredentials()
	client, err := rocket.NewClient(nil, rocket.ClientOptions{
		ServerUrl:   serverUrl,
		Credentials: credentials,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}
