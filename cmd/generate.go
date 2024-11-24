package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Weburz/repoforge/internal/template"
)

var generateCmdShortUsage = "Scaffold a project from a template"

var generateCmdLongUsage = `
Scaffold a project from a template.

Use this command to scaffold a project from a template stored either locally or
in a remote location (like a GitHub/GitLab repository). Support for other remote
storage environments will be supported in a future version.
`

var generateCmdExample = "forge generate \"Weburz/nuxt-base\""

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

func init() {
	rootCmd.AddCommand(templateCmd)
}
