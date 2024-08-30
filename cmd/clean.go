package cmd

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean/delete all downloaded templates.",
	Run: func(cmd *cobra.Command, args []string) {
		templates_dir := fmt.Sprintf("%s/repoforge", xdg.DataHome)
		templates, err := os.ReadDir(templates_dir)

		if err != nil {
			fmt.Printf("Failed to find any templates at %s\n", templates)
		}

		for _, template := range templates {
			// TODO: Implement the functionality to remove the directories
			fmt.Printf("Deleting %s\n", template.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
