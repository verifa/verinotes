// Copyright © 2023 Verifa <info@verifa.io>
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
Config file would be useful to specify custom port etc., but perhaps envconfig
is preferred.
var (
	configFiles []string
)
*/

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "verinotes",
	Short: "VeriNotes - for note taking",
	Long: `VeriNotes

At the moment there's only one command you should be interested in:

./verinotes server

This brings up the VeriNotes application listening on port 3000, the UI
is embedded under path /ui, but users will be automatically forwarded to it.
The API is available under /api/v1.`,
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

/*
func init() {
	rootCmd.PersistentFlags().StringSliceVar(&configFiles, "config", nil, "config files to parse")
}
*/
