package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/teakingwang/cursor-demo/cmd/app"
)

var (
	port string
)

var rootCmd = &cobra.Command{
	Use:   "cursor-demo",
	Short: "A demo application using Gin framework",
	Long:  `A demo application using Gin framework with complete features.`,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Long:  `Start the HTTP server with specified port`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server := app.NewServer(port)
		return server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "端口号")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
