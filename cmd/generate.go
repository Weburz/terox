package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a template locally for future development needs.",
	Aliases: []string{"gen"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf(
				"Command accepts 0 arguments but %d were passed",
				len(args),
			)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Downloading templates...")
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
