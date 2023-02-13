/*
Copyright © 2022 Jacob Larfors <jlarfors@verifa.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/verifa/verinotes/server"
)

// var serverConfig server.Config

// var storeConfig store.Config

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.TODO()

		// TODO see if needed
		//if serverConfig.DevMode {
		//}

		srv, err := server.New(ctx)
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
	// envErr := envconfig.Process("CL", &serverConfig)
	// cobra.CheckErr(envErr)

	rootCmd.AddCommand(serverCmd)

	// serverCmd.Flags().BoolVarP(&serverConfig.DevMode, "dev", "d", serverConfig.DevMode, "Enable dev mode")
}