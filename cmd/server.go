// Copyright © 2023 Verifa <info@verifa.io>
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"github.com/verifa/verinotes/server"
	"github.com/verifa/verinotes/store"
)

// var serverConfig server.Config

var storeConfig store.Config

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts VeriNotes server",
	Long:  `Starts the VeriNotes server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.TODO()

		store, err := store.New(ctx, &storeConfig)
		if err != nil {
			return fmt.Errorf("creating store: %w", err)
		}

		srv, err := server.New(ctx, store)
		if err != nil {
			return fmt.Errorf("creating server: %w", err)
		}

		addr := ":3000"
		l, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("listening at %s: %w", addr, err)
		}

		log.Println("verinotes listening on", addr)

		return http.Serve(l, srv)
	},
}

func init() {
	// Handle environment variable configs before parsing command line args
	envErr := envconfig.Process("VN", &storeConfig)
	cobra.CheckErr(envErr)

	rootCmd.AddCommand(serverCmd)
}
