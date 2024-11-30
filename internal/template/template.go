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
