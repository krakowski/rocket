package cmd

import (
	"bytes"
	"fmt"
	rocket "github.com/krakowski/rocket/api"
	"github.com/krakowski/rocket/util"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	defaultDirectory = ".rocket"
)

var (
	directory string
)

type TemplateValues struct {
	Env map[string]string
}

var notifyCommand = &cobra.Command{
	Use:           "notify TEMPLATE",
	Short:         "Sends a notification to a channel",
	SilenceErrors: true,
	Args:          cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		spin := util.StartSpinner("Parsing template file")

		// Parse the template file
		tmpl, err := template.ParseFiles(path.Join(directory, args[0]+".yml"))
		if err != nil {
			spin.StopError(err)
			os.Exit(1)
		}

		// Read environment variables
		values, err := createTemplateValues()
		if err != nil {
			spin.StopError(err)
			os.Exit(1)
		}

		// Execute the template
		var out bytes.Buffer
		if err = tmpl.Execute(&out, values); err != nil {
			spin.StopError(err)
			os.Exit(1)
		}

		// Decode the yaml output
		var payload rocket.MessagePayload
		if err = yaml.NewDecoder(bytes.NewReader(out.Bytes())).Decode(&payload); err != nil {
			spin.StopError(err)
			os.Exit(1)
		}

		spin.StopSuccess("Success")

		spin = util.StartSpinner("Authenticating with RocketChat")

		// Create a new RocketChat client
		client, err := util.NewRocketClient()
		if err != nil {
			spin.StopError(err)
			fmt.Printf(" - %s\n", rocket.LastError.Message)
			os.Exit(1)
		}

		spin.StopSuccess("Logged in as " + client.CurrentUser)

		spin = util.StartSpinner("Sending notification")

		// Post the message
		resp, err := client.Message.Post(payload)
		if err != nil {
			spin.StopError(err)
			fmt.Printf(" - %s\n", rocket.LastError.Message)
			os.Exit(1)
		}

		spin.StopSuccess("Sent to channel " + resp.Channel)
	},
}

func createTemplateValues() (TemplateValues, error) {
	envMap := make(map[string]string)
	for _, v := range os.Environ() {
		pair := strings.Split(v, "=")
		envMap[pair[0]] = pair[1]
	}

	return TemplateValues{Env: envMap}, nil
}

func init() {
	notifyCommand.Flags().StringVarP(&directory, "directory", "d", defaultDirectory, "The directory in which to search for templates")
}
