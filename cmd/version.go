// Copyright © 2023 Verifa <info@verifa.io>
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "dev"
	date    = "dev"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the VeriNotes version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("VeriNotes version", version)
		fmt.Println("Commit:", commit)
		fmt.Println("Date:", date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
