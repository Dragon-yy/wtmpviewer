package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"wtmpviewer/internal/logparser"

	"github.com/spf13/cobra"
)

// checkSecureCmd represents the check-secure command
var checkSecureCmd = &cobra.Command{
	Use:   "check-secure",
	Short: "Check secure log files for successful logins",
	Long:  `This command parses all secure log files in a specified directory and prints successful login attempts.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the directory flag
		dir, err := cmd.Flags().GetString("directory")
		if err != nil || dir == "" {
			fmt.Println("Please provide a valid directory using --directory")
			return
		}

		// Walk through the directory
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Check if the file name starts with "secure"
			if !info.IsDir() && strings.HasPrefix(filepath.Base(path), "secure") {
				fmt.Printf("Processing file: %s\n", path)
				records, err := logparser.ParseSecure(path)
				if err != nil {
					fmt.Printf("Error reading secure file %s: %v\n", path, err)
					return nil
				}

				// Print all parsed records
				for _, record := range records {
					fmt.Println(record)
				}
			}
			return nil
		})

		// Handle any errors from filepath.Walk
		if err != nil {
			fmt.Printf("Error walking the path %s: %v\n", dir, err)
		}
	},
}

func init() {
	// Register the command with the root command
	rootCmd.AddCommand(checkSecureCmd)
	// Add a flag for the directory to scan
	checkSecureCmd.Flags().StringP("directory", "d", "/var/log", "Directory to scan for secure files")
}
