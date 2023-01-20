package cmd

import (
	"github.com/mheers/dex-local-discovery/helpers"
	"github.com/spf13/cobra"
)

var (
	// LogLevelFlag describes the verbosity of logs
	LogLevelFlag string
	// ConfigFileFlag holds the path to the config file
	ConfigFileFlag string

	// OutputFormatFlag can be json, yaml or table
	OutputFormatFlag string

	// // Config holds the read config
	// Config *config.Config

	rootCmd = &cobra.Command{
		Use:   "dex-local-discovery",
		Short: "dex-local-discovery is a sidecar for dex to rewrite the openid-configuration on the fly",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			helpers.PrintInfo()
			cmd.Help()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&LogLevelFlag, "log-level", "l", "info", "possible values are debug, error, fatal, panic, info, trace")
	rootCmd.PersistentFlags().StringVarP(&OutputFormatFlag, "output-format", "O", "table", "format [json|table|yaml|csv]")
}
