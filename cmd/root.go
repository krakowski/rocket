package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

const (
	version = "1.0.1"
)

var rootCommand = &cobra.Command{
	Use:           "rocket",
	Short:         "A simple command line interface for sending notifications to RocketChat",
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCommand.AddCommand(notifyCommand)
	rootCommand.Version = version
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
