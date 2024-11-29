/**
 * Package cmd - The "cmd" package contains the logic to handle the various
 * commands of the CLI application.
 *
 * The "generate" file in particular, contained in the "cmd" package handles
 * the logic to download (if needed) and scaffold a project from a pre-existing
 * template.
 */
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Weburz/terox/internal/template"
)

var generateCmdShortUsage = "Scaffold a project from a template"

var generateCmdLongUsage = `
Scaffold a project from a template.

Use this command to scaffold a project from a template stored either locally or
in a remote location (like a GitHub/GitLab repository). Support for other remote
storage environments will be supported in a future version.
`

var generateCmdExample = "terox generate \"Weburz/nuxt-base\""

// Handle the logic for the "generate" command
var templateCmd = &cobra.Command{
	Use:     "generate",
	Short:   generateCmdShortUsage,
	Long:    generateCmdLongUsage,
	Aliases: []string{"gen"},
	Example: generateCmdExample,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create the template
		tg, err := template.NewTemplateGenerator(args[0])

		// Throw an error if the template generation failed for whatever reason
		if err != nil {
			rootCmd.PrintErrf("Error: %v\n", err)
			return
		}

		// Scaffold a project using the template created above
		if err := tg.Scaffold(); err != nil {
			rootCmd.PrintErrf("Error: %v\n", err)
		}
	},
}

// Register the logic for the "generate" command to the root CLI application
func init() {
	rootCmd.AddCommand(templateCmd)
}
