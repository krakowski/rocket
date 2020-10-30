package util

import (
	"fmt"
	"github.com/krakowski/rocket/api"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"syscall"
)

const (
	envUser = "ROCKET_USER"
	envPassword = "ROCKET_PASS"
)

func GetCredentials() rocket.Credentials {

	// Get the username
	username, present := os.LookupEnv(envUser)
	if present == false {
		_, _ = fmt.Fprint(os.Stderr, "RocketChat username: ")
		_, err := fmt.Scanln(&username)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Get the password
	password, present := os.LookupEnv(envPassword)
	if present == false {
		_, _ = fmt.Fprint(os.Stderr, "RocketChat password: ")
		inputBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		password = string(inputBytes)
		_, _ = fmt.Fprintln(os.Stderr)
	}

	return rocket.Credentials{
		Username: username,
		Password: password,
	}
}
