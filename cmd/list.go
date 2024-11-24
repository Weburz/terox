/**
 * Package cmd - The "cmd" package contains the logic to handle various
 * commands passed to the CLI tool.
 *
 * The "list" file in particular contains the logic to list all locally
 * available template(s) on disk.
 */
package cmd

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

// Logic to handle the "list" command
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

// Register the logic above with the CLI application
func init() {
	rootCmd.AddCommand(listCmd)
}

/**
 * listTemplates - List all available templates on disk.
 *
 * Parameters:
 * None
 *
 * Returns:
 * None
 */
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
