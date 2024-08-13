package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"wtmpviewer/internal/logparser/wtmp"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the wtmp logs",
	Long:  `This command allows you to view the contents of the wtmp log file.`,
	Run: func(cmd *cobra.Command, args []string) {
		wtmpFile, err := cmd.Flags().GetString("file")
		if err != nil || wtmpFile == "" {
			fmt.Println("Please provide a valid wtmp file using --file")
			return
		}

		err = wtmp.ParseWtmp2(wtmpFile)
		if err != nil {
			fmt.Printf("Error reading wtmp file: %v\n", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().StringP("file", "f", "/var/log/wtmp", "Path to the wtmp file")
}
