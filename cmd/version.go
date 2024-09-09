package cmd

import (
	"github.com/spf13/cobra"
)

type Version struct {
	Version   string
	Commit    string
	BuildTime string
}

var version *Version

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number and exit (also --version)",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("forge %s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func SetVersion(v *Version) {
	version = v
}
