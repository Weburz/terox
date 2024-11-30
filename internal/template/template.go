package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

/**
 * ListTemplates: List all locally available templates.
 *
 * Parameters:
 *   None
 *
 * Returns:
 *   None
 */
func ListTemplates() error {
	// The filepath where the templates are "usually" stored
	templatesDir := filepath.Join(xdg.DataHome, "terox")

	// Read the contents of the template directory to memory
	templates, err := os.ReadDir(templatesDir)

	// Throw an error and exit with non-zero code if the directory was not found
	if err != nil {
		return fmt.Errorf("Failed to read contents of %s: %s\n", templates, err)
	}

	// If the templates were found then list them to STDOUT
	if len(templates) != 0 {
		fmt.Printf("Available Templates:\n")

		for _, template := range templates {
			if template.IsDir() {
				fmt.Printf("%s\n", template.Name())
			}
		}
	} else {
		fmt.Printf("Available Templates: None\n")
	}

	return nil
}

/**
 * Clean: Cleanup the locally downloaded templates.
 *
 * Parameters:
 *   None
 *
 * Returns:
 *   None
 */
func Clean() error {
	// Get the path to the template directory
	templatesDir := filepath.Join(xdg.DataHome, "terox")

	// Read the contents of the template directory to check for templates
	templates, err := os.ReadDir(templatesDir)

	// Throw an error if the template directory is unreachable
	if err != nil {
		return fmt.Errorf(
			"Failed to find any templates at %s: %w",
			templates,
			err,
		)
	}

	// Prompt the user about what template deletion
	fmt.Printf("The following templates were deleted:\n\n")

	// If templates were found locally, then attempt to delete them else throw
	// an error
	if len(templates) != 0 {
		for _, template := range templates {
			templatePath := filepath.Join(templatesDir, template.Name())
			fmt.Printf("%s\n", template.Name())

			if err := os.RemoveAll(templatePath); err != nil {
				return fmt.Errorf(
					"Failed to delete %s: %w",
					template.Name(),
					err,
				)
			}
		}
	}

	return nil
}
