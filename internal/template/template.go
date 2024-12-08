package template

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

// The base directory where the templates will be stored locally
var templateDir = filepath.Join(xdg.DataHome, "terox")

/**
 * Template - The struct to store the information related to a template.
 *
 * Fields:
 * Path (string): The path on the local system where the template will be
 *    downloaded from.
 */
type Template struct {
	Path string
}

/**
 * FunctionName - Create an instance of the "Template" struct and return it (or
 * throw an error when needed).
 *
 * Parameters:
 * repo (string): The GitHub repository URL the template is located at.
 *
 * Returns:
 * Returns a pointer to the "Template" struct or an error
 */
func NewTemplate(repo string) (*Template, error) {
	// Split the GitHub repository URL at the '/' for further processing
	parts := strings.Split(repo, "/")

	// Throw an error if the passed URL is not of the form - "Weburz/terox"
	if len(parts) != 2 {
		return nil,
			fmt.Errorf("Invalid repository format, expected \"<OWNER>/<REPO>\"")
	}

	// Name of the template to store locally
	name := parts[1]

	// Return an instance of the "Template" struct (or throw an error, if any)
	return &Template{
		Path: filepath.Join(templateDir, name),
	}, nil
}

/**
 * Scaffold - Scaffold the project by download the template (if necessary)
 *
 * Parameters:
 * None
 *
 * Returns:
 * A wrapped error if any is raised by the underlying wrapping function.
 */
func (t *Template) Scaffold() error {
	// Check if the template already exists locally
	if _, err := os.Stat(t.Path); os.IsNotExist(err) {
		fmt.Printf("Template not found locally...downloading\n")
		return nil
	} else if err != nil {
		return fmt.Errorf("Error checking the template path: %w", err)
	}

	fmt.Printf("Template found locally at: %s\n", t.Path)

	return nil
}

/**
 * List - List all the locally available templates.
 *
 * Parameters:
 * None
 *
 * Returns:
 * A wrapped error if any is raised.
 */
func List() error {
	// Check if any template exists locally, if yes, list them to STDOUT
	if templates, err := os.ReadDir(templateDir); err != nil {
		return fmt.Errorf(
			"Failed to read the contents of %s directory: %w",
			templateDir,
			err,
		)
	} else if len(templates) != 0 {
		fmt.Printf("Available Templates:\n")
		for _, template := range templates {
			if template.IsDir() {
				fmt.Printf("%s\n", template.Name())
			}
		}
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
 *   An error if any was raised during the removal process.
 */
func Clean() error {
	// Read the contents of the template directory to check for templates
	if templates, err := os.ReadDir(templateDir); err != nil {
		return fmt.Errorf(
			"Failed to find any templates at %s: %w",
			templateDir,
			err,
		)
	} else if len(templates) != 0 {
		fmt.Printf("The following templates were deleted:\n\n")
		for _, template := range templates {
			path := filepath.Join(templateDir, template.Name())
			fmt.Printf("%s\n", template.Name())
			if err := os.RemoveAll(path); err != nil {
				return fmt.Errorf(
					"Failed to remove %s: %w",
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
