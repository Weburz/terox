package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// The location on disk to create the template to
var location string

var shortUsage = "Create a project template"

var longUsage = `
Create a project template.

Use this command to create a project template which can later be used to
scaffold future projects! By default, the project is generated under the current
working directory but it is possible to configure it by passing the "--location"
flag.
`

var example = "forge create \"simple-website\" --location=\".\""

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   shortUsage,
	Long:    longUsage,
	Example: example,
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
