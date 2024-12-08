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

var scaffoldCmdShortHelp = "Scaffold a project from a template"

var scaffoldCmdLongHelp = `
Scaffold a project from a template.

Use this command to scaffold a project from a template stored either locally or
in a remote location (like a GitHub/GitLab repository). Support for other remote
storage environments will be supported in a future version.
`

var generateCmdExample = "terox generate \"Weburz/nuxt-base\""

// Handle the logic for the "generate" command
var scaffoldCmd = &cobra.Command{
	Use:     "scaffold",
	Short:   scaffoldCmdShortHelp,
	Long:    scaffoldCmdLongHelp,
	Example: generateCmdExample,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new template instance
		t, err := template.NewTemplate(args[0])

		// Throw an error if any was raised when creating a template instance
		if err != nil {
			rootCmd.PrintErrf("Error: %v\n", err)
		}

		// Attempt to scaffold the project based on the template else throw an
		// error and exit the execution sequence.
		if err := t.Scaffold(args[0]); err != nil {
			rootCmd.PrintErrf("Error: %v\n", err)
		}
	},
}

// Register the logic for the "generate" command to the root CLI application
func init() {
	rootCmd.AddCommand(scaffoldCmd)
}
