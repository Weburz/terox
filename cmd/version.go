package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const version = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("RepoForge Build Information\n")
		fmt.Printf("\nVersion: %s\n", version)
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
