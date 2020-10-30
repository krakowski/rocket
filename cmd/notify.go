package cmd

import (
	"bytes"
	rocket "github.com/krakowski/rocket/api"
	"github.com/krakowski/rocket/util"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
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

		// Parse the template file
		tmpl, err := template.ParseFiles(path.Join(directory, args[0]+".yml"))
		if err != nil {
			log.Fatal(err)
		}

		// Read environment variables
		values, err := createTemplateValues()
		if err != nil {
			log.Fatal(err)
		}

		// Execute the template
		var out bytes.Buffer
		if err = tmpl.Execute(&out, values); err != nil {
			log.Fatal(err)
		}

		// Decode the yaml output
		var payload rocket.MessagePayload
		if err = yaml.NewDecoder(bytes.NewReader(out.Bytes())).Decode(&payload); err != nil {
			log.Fatal(err)
		}

		// Create a new RocketChat client
		client := util.NewRocketClient()

		// Post the message
		if err = client.Message.Post(payload); err != nil {
			log.Fatal(err)
		}
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
