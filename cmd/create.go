/**
 * Package cmd - The "cmd" package consists the logic to handle the commands
 * passed to the CLI tool.
 *
 * The "create" file of the "cmd" package in particular is responsible for
 * creating a fresh new template for future use (as in scaffolding projects
 * from it).
 */
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

// Logic to handle the "create" command for the CLI tool.
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

// Register the "create" command for the CLI tool and also add a flag which
// accepts a filepath to create the template at (defaults to the current working
// directory).
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

/**
 * createTemplate - Create a template (with a give name) and to a specified
 *     path.
 *
 * Parameters:
 * name string: The name to assign to the template.
 * templatePath: The path to create (and store) the template files at. Defaults
 *     to the current working directory.
 *
 * Returns:
 * None
 */
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

	jsonContent := []byte("{}\n")

	// Add some basic necessary contents to the "forge.json" file
	if err := os.WriteFile(
		filepath.Join(templatePath, "forge.json"),
		jsonContent,
		0644,
	); err != nil {
		rootCmd.PrintErrf("Failed to write to \"forge.json\" due to %s", err)
	}
}
