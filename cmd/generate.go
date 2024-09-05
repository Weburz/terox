package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Download a template from a remote source.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Downloading templates...")
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
