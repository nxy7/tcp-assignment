package cmd

import (
	"fmt"
	"os"

	"github.com/nxy7/tcp-assignment/cmd/client"
	"github.com/nxy7/tcp-assignment/cmd/serve"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	// Use:   "tcp",
	Short: "Custom TCP protocol",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serve.ServeCmd)
	rootCmd.AddCommand(client.ClientCmd)
}
