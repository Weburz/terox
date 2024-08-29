package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all locally available templates.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: All available templates")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
