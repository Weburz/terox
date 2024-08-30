package cmd

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all locally available templates.",
	Run: func(cmd *cobra.Command, args []string) {
		templates_dir := fmt.Sprintf("%s/repoforge", xdg.DataHome)
		templates, err := os.ReadDir(templates_dir)

		if err != nil {
			fmt.Printf("No templates found at %s\n", templates_dir)
			os.Exit(1)
		}

		fmt.Printf("Templates directory: %s\n\n", templates_dir)
		fmt.Printf("Available templates:\n")

		for _, template := range templates {
			if template.IsDir() {
				fmt.Printf("%s\n", template.Name())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
