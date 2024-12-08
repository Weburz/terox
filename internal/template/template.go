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

/**
 * Template - The struct to store the information related to a template.
 *
 * Fields:
 * Repo (string): The GitHub URL to fetch the template repository from.
 * Path (string): The path on the local system where the template will be
 *    downloaded from.
 */
type Template struct {
	TemplatePath string
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

	// Create a variable to store the absolute path of the template
	templatePath, _ := filepath.Abs(
		filepath.Join(xdg.DataHome, "terox", parts[1]),
	)

	// Return an instance of the "Template" struct (or throw an error, if any)
	return &Template{
		TemplatePath: templatePath,
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
	if _, err := os.Stat(t.TemplatePath); os.IsNotExist(err) {
		fmt.Printf("Template not found locally...downloading\n")
		return nil
	} else if err != nil {
		return fmt.Errorf("Error checking the template path: %w", err)
	}

	fmt.Printf("Template found locally at: %s\n", t.TemplatePath)

	return nil
}

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
