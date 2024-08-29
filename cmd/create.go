package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a template locally.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating template...")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
