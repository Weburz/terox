package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// The location on disk to create the template to
var location string

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a project template",
	Example: `forge create "simple-website" --location="."`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		absLocation, _ := filepath.Abs(location)
		createTemplate(args[0], absLocation)
	},
}

func init() {
	// Fetch the current working directory
	cwd, _ := os.Getwd()
	defaultLocation, _ := filepath.Abs(cwd)

	// Add a flag called "--location" whose default value is the current working
	// directory
	createCmd.Flags().StringVar(
		&location,
		"location",
		defaultLocation,
		"specify the location to create the template at (e.g., ./templates)",
	)

	rootCmd.AddCommand(createCmd)
}

func createTemplate(name string, absLocation string) {
	rootCmd.Printf("Creating the template \"%s\" at %s\n", name, absLocation)
}
