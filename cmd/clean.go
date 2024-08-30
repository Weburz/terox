package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"gc", "cleanup"},
	Short:   "Clean/delete all downloaded templates.",
	Example: "forge clean\nforge gc\nforge cleanup",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("Command does not accept any arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		templatesDir := filepath.Join(xdg.DataHome, "repoforge")
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
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
