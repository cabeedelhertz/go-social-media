package main

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "social",
	}
	flags := rootCmd.PersistentFlags()
	flags.String("service_name", "social", "the name the service")
	flags.Int("system_port", 9102, "the system port to listen on")
	rootCmd.AddCommand(NewServerCommand())

	return rootCmd
}

func main() {
	if err := NewRootCommand().Execute(); err != nil {
		os.Exit(2)
	}
}
