package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all locally available templates.",
	Run: func(cmd *cobra.Command, args []string) {
		templates_dir := filepath.Join(xdg.DataHome, "repoforge")
		templates, err := os.ReadDir(templates_dir)

		// Throw error and exit execution if the data directory is unreadable
		if err != nil {
			fmt.Printf("Failed to read %s directory: %v\n", templates_dir, err)
			os.Exit(1)
		}

		fmt.Printf("Templates directory: %s\n\n", templates_dir)

		// Share conditional message to user if no templates were found
		if len(templates) == 0 {
			fmt.Printf("Available templates: None\n")
		} else {
			fmt.Printf("Available templates:\n")
			for _, template := range templates {
				if template.IsDir() {
					fmt.Printf("%s\n", template.Name())
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
