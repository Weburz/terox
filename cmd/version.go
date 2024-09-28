package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number and exit",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("RepoForge %s\n", "v0.0.1-alpha")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
