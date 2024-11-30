package template

import (
	"archive/zip"
	"fmt"
	"io"
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

/**
 * Extract: Extract a template's zipball to a specified destination.
 *
 * Parameters:
 *   zipfile: (string) The filepath to the downloaded zipfile.
 *   dest: (string) The destination path to extract the zipped contents to.
 *
 * Returns:
 *   An error message (if any).
 */
func Extract(zipfile, templateDest string) error {
	// Set the destination to extract the zipfile contents to
	// Should default to ~/.local/share/terox
	dest := filepath.Join(xdg.DataHome, "terox", templateDest)

	// Read the zipfile and close it when the function completes execution
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		return fmt.Errorf("Failed to open zip file: %w", err)
	}
	defer r.Close()

	// Ensure the destination directory exists
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create destination directory: %w", err)
	}

	// Iterate through the contents of the zipfile
	for _, f := range r.File {
		filePath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, f.Mode()); err != nil {
				return fmt.Errorf("Failed to create directory: %w", err)
			}
			continue
		}

		srcFile, err := f.Open()
		if err != nil {
			return fmt.Errorf("Failed to open file in archive: %w", err)
		}
		defer srcFile.Close()

		destFile, err := os.Create(dest)
		if err != nil {
			return fmt.Errorf("Failed to create file: %w", err)
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, srcFile); err != nil {
			return fmt.Errorf("Failed to copy file contents: %w", err)
		}
	}

	// TODO: Implement functionality to remove the downloaded zipfile if its
	// extraction was a success

	return nil
}
