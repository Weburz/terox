package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var generateCmdShortUsage = "Scaffold a project from a template"

var generateCmdLongUsage = `
Scaffold a project from a template.

Use this command to scaffold a project from a template stored either locally or
in a remote location (like a GitHub/GitLab repository). Support for other remote
storage environments will be supported in a future version.
`

var generateCmdExample = "forge generate \"Weburz/nuxt-base\""

var templateCmd = &cobra.Command{
	Use:     "generate",
	Short:   generateCmdShortUsage,
	Long:    generateCmdLongUsage,
	Aliases: []string{"gen"},
	Example: generateCmdExample,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create a variable representation of the path to the template
		templatePath, _ := filepath.Abs(
			filepath.Join(
				xdg.DataHome,
				"repoforge",
				// Represents the GitHub/GitLab user/organisation like "Weburz"
				strings.Split(args[0], "/")[0],
				// Represents the GitHub/GitLab repository like "nuxt-base"
				strings.Split(args[0], "/")[1],
			),
		)

		// Handle the logic to scaffold the project if the template is already
		// available locally, else download it from the remote location
		_, err := os.Stat(templatePath)
		if errors.Is(err, os.ErrNotExist) {
			rootCmd.Printf("No template named '%s' found locally\n", args[0])
			rootCmd.Printf("Downloading template to %s\n", templatePath)

			os.MkdirAll(templatePath, os.ModePerm)

			url := fmt.Sprintf(
				"https://api.github.com/repos/%s/zipball",
				args[0],
			)

			resp, err := http.Get(url)

			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()

			file, _ := os.Create("template.zip")
			defer file.Close()

			io.Copy(file, resp.Body)
		} else {
			rootCmd.Printf("Scaffolding project from %s\n", templatePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
