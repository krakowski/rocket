package util

import (
	rocket "github.com/krakowski/rocket/api"
	"log"
	"os"
)

const (
	envHost = "ROCKET_HOST"
)

func NewRocketClient() (*rocket.Client) {
	serverUrl, present := os.LookupEnv(envHost)
	if !present {
		log.Fatal("Please specify the RocketChat server url in the ROCKET_HOST environment variable")
	}

	credentials := GetCredentials()
	client, err := rocket.NewClient(nil, rocket.ClientOptions{
		ServerUrl:   serverUrl,
		Credentials: credentials,
	})

	if err != nil {
		log.Fatal(err)
	}

	return client
}
