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

			// Create the directory where the contents of the template will be
			// stored locally
			os.MkdirAll(templatePath, os.ModePerm)

			// Create a temporary file for the zipball contents
			tempfile, err := os.CreateTemp("", "repoforge-*.zip")
			if err != nil {
				rootCmd.PrintErrf("Failed to create temp file: %v\n", err)
			}
			defer os.Remove(tempfile.Name())

			// Download the zipball built in the previous command above
			if err := downloadTemplate(args[0], tempfile.Name()); err != nil {
				rootCmd.PrintErrf("Failed to download zipball: %v\n", err)
			}
		} else {
			rootCmd.Printf("Scaffolding project from %s\n", templatePath)
		}
	},
}

// Download a zipped archive of the repository from GitHub and store it in a
// temporary directory for extracting its contents
func downloadTemplate(repo, filepath string) error {
	// Set the API endpoint to fetch the zipball from
	url := fmt.Sprintf("https://api.github.com/repos/%s/zipball", repo)

	// Fetch the zipball using a HTTP GET request to the API endpoint above.
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to download template: %w\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad server response: %d\n", resp.StatusCode)
	}

	// Create a zipball in the temporary directory created above
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Failed to create zipball: %w\n", err)
	}
	defer file.Close()

	// Copy the contents returned from the GET request to the zipball onto the
	// local filesystem.
	_, err = io.Copy(file, resp.Body)

	// Return an error if the GET request from GitHub failed to be converted
	// into a zipball
	return err
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
