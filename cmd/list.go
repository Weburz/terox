package cmd

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "show"},
	Short:   "List all locally available templates.",
	Example: "forge list\nforge ls\nforge show",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		listTemplates()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTemplates() {
	templates_dir := filepath.Join(xdg.DataHome, "repoforge")
	templates, err := os.ReadDir(templates_dir)

	// Throw error and exit execution if the data directory is unreadable
	if err != nil {
		rootCmd.PrintErrf(
			"Failed to read %s directory: %v\n",
			templates_dir,
			err,
		)
		os.Exit(1)
	}

	// Print the directory where the templates are located at
	rootCmd.Printf("Templates directory: %s\n\n", templates_dir)

	// Share conditional message to user if no templates were found
	if len(templates) == 0 {
		rootCmd.Printf("Available templates: None\n")
	} else {
		rootCmd.Printf("Available templates:\n")
		for _, template := range templates {
			if template.IsDir() {
				rootCmd.Printf("%s\n", template.Name())
			}
		}
	}
}
