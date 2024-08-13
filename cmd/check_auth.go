package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"wtmpviewer/internal/logparser"

	"github.com/spf13/cobra"
)

// checkAuthCmd represents the check-auth command
var checkAuthCmd = &cobra.Command{
	Use:   "check-auth",
	Short: "Check auth log files for successful logins",
	Long:  `This command parses all auth log files in a specified directory and prints successful login attempts.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.Flags().GetString("directory")
		if err != nil || dir == "" {
			fmt.Println("Please provide a valid directory using --directory")
			return
		}

		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && filepath.Base(path) == "auth.log" {
				fmt.Printf("Processing file: %s\n", path)
				records, err := logparser.ParseAuthLog(path)
				if err != nil {
					fmt.Printf("Error reading auth log file %s: %v\n", path, err)
					return nil
				}

				for _, record := range records {
					fmt.Println(record)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking the path %s: %v\n", dir, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkAuthCmd)
	checkAuthCmd.Flags().StringP("directory", "d", "/var/log", "Directory to scan for auth.log files")
}
