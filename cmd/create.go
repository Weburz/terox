package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// The templatePath on disk to create the template to
var templatePath string

var shortUsage = "Create a project template"

var longUsage = `
Create a project template.

Use this command to create a project template which can later be used to
scaffold future projects! By default, the project is generated under the current
working directory but it is possible to configure it by passing the "--location"
flag.
`

var example = "forge create \"simple-website\" --path=\".\""

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   shortUsage,
	Long:    longUsage,
	Example: example,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the template directory
		templatePath, err := filepath.Abs(filepath.Join(templatePath, args[0]))

		// Warn user and exit execution if an error was raised
		if err != nil {
			rootCmd.PrintErrf(
				"Failed to identify template directory at %s\n",
				templatePath,
			)
			os.Exit(1)
		}

		// Create the template
		createTemplate(args[0], templatePath)
	},
}

func init() {
	// Fetch the current working directory
	cwd, _ := os.Getwd()

	// The filepath where the template directory will be created by default
	defaultPath, _ := filepath.Abs(cwd)

	// Add a flag called "--location" whose default value is the current working
	// directory
	createCmd.Flags().StringVar(
		&templatePath,
		"path",
		defaultPath,
		"specify the path to create the template at (e.g., ./templates)",
	)

	rootCmd.AddCommand(createCmd)
}

func createTemplate(name string, templatePath string) {
	rootCmd.Printf("Creating the template \"%s\" at %s\n", name, templatePath)

	// Create the template directory
	if err := os.MkdirAll(templatePath, os.ModePerm); err != nil {
		rootCmd.PrintErrf(
			"Failed to create  template directory at %s\n",
			templatePath,
		)
		panic(err)
	}

	// Create the JSON "data" file to fetch some local information about the
	// template from
	if _, err := os.Create(
		filepath.Join(templatePath, "forge.json"),
	); err != nil {
		rootCmd.PrintErrf(
			"Failed to write the forge.json file to %s\n",
			templatePath,
		)
		panic(err)
	}

}
