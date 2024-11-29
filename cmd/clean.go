/**
 * Package cmd - The "cmd" package contains the CLI commands for the tool.
 *
 * The "clean" file in particular which is part of the "cmd" package contains
 * the logic to clean up downloaded template(s) on disk.
 */
package cmd

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

// Command to handle the "clean" command of the CLI tool
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"gc", "cleanup"},
	Short:   "Clean/delete all downloaded templates.",
	Example: "terox clean\nterox gc\nterox cleanup",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cleanTemplates()
	},
}

// Register the command to the CLI tool
func init() {
	rootCmd.AddCommand(cleanCmd)
}

/**
 * cleanTemplates - Cleanup the template(s) from disk.
 *
 * Parameters:
 * None
 *
 * Returns:
 * None
 */
func cleanTemplates() {
	templatesDir := filepath.Join(xdg.DataHome, "terox")
	templates, err := os.ReadDir(templatesDir)

	// Throw error and exit execution loop if no templates were found
	if err != nil {
		rootCmd.Printf("Failed to find any templates at %s\n", templates)
		os.Exit(1)
	}

	// Warn user and exit execution loop safely if no templates were found
	if len(templates) == 0 {
		rootCmd.Println("No templates were found.")
		os.Exit(0)
	}

	// Print the location of the directory where the templates are stored at
	rootCmd.Printf("The following templates were deleted:\n\n")

	// Loop through the templates directory and delete everything
	// (including files and folders) inside it
	for _, template := range templates {
		templatePath := filepath.Join(templatesDir, template.Name())

		rootCmd.Printf("%s\n", template.Name())

		err := os.RemoveAll(templatePath)

		if err != nil {
			rootCmd.Printf(
				"Failed to delete %s: %v\n",
				template.Name(),
				err,
			)
			os.Exit(1)
		}
	}
}
