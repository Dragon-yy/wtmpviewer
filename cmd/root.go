package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "wtmp-viewer",
	Short: "A tool to view wtmp logs",
	Long:  `wtmp-viewer is a CLI tool to view and parse Linux wtmp logs.`,
}

// Execute runs the root command, setting up the Cobra command structure.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define flags and configuration settings for the root command.
	// Persistent flags are global and available for all subcommands.
	// rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.wtmp-viewer.yaml)")
}
