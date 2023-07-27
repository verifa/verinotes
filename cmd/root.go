// Copyright © 2023 Verifa <info@verifa.io>
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	configFiles []string
	// dbConfig
	// notesConfig
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github.com/verifa/verinotes",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error executing the CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// TODO envconfig, or no envconfig :shrug:
	// envErr := envconfig.Process("CL", &requestsConfig)
	// cobra.CheckErr(envErr)

	rootCmd.PersistentFlags().StringSliceVar(&configFiles, "config", nil, "config files to parse")
}

// initConfig reads in .env config files, if any
func initConfig() {
	// If configFiles is nil, godotenv will look for a local .env file by default, which
	// is kind of unexpected for a user
	if configFiles != nil {
		err := godotenv.Overload(configFiles...)
		cobra.CheckErr(err)
	}
}
