/**
 * Package cmd - The "cmd" package contains the logic to handle the commands
 * passed to the CLI tool.
 *
 * The "version" file of the "cmd" package contains the simple to print out
 * relevant version information of the CLI application.
 */
package cmd

import (
	"github.com/spf13/cobra"
)

// Handle the logic for the "version" command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number and exit",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("RepoForge %s\n", "v0.0.1-alpha")
	},
}

// Register the "version" command for the CLI tool
func init() {
	rootCmd.AddCommand(versionCmd)
}
