/**
 * Package template - The package contains the logic to download and scaffold
 * the project. This file in particular which is part of the package contains
 * the core logic to create the objects necessary to parse the arguments passed
 * to the "generate" CLI and using them to fetch/scaffold the project.
 */
package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

/**
 * TemplateGenerator - A struct to represent the template's GitHub repository
 * and the filepath it should be stored at.
 *
 * Fields:
 * Owner: string // The user/organisation who owns the template's repository.
 * Repo: string // The template's repository (on GitHub).
 * TemplatePath: string // The path (on disk) where the template should be
 *		stored at.
 */
type TemplateGenerator struct {
	Owner        string
	Repo         string
	TemplatePath string
}

/**
 * NewTemplateGenerator - The factory function to create an instance of the
 *     "TemplateGenerator" struct.
 *
 * Parameters:
 * repo string // The repository name in the form: "Weburz/repoforge".
 *
 * Returns:
 * Returns an instance of the "TemplateGenerator" struct.
 */
func NewTemplateGenerator(repo string) (*TemplateGenerator, error) {
	// Get the repository and owner name from the URL
	parts := strings.Split(repo, "/")

	if len(parts) != 2 {
		return nil,
			fmt.Errorf(
				"Invalid repository format, expected \"<OWNER>/<REPO>\"",
			)
	}

	// Create a filepath according to the XDG Base Directory specification
	templatePath, err := filepath.Abs(
		filepath.Join(xdg.DataHome, "repoforge", parts[0], parts[1]),
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to determine template path: %w", err)
	}

	return &TemplateGenerator{
		Owner:        parts[0],
		Repo:         parts[1],
		TemplatePath: templatePath,
	}, nil
}

/**
 * Scaffold - Scaffold the project (by downloading the template, if necessary)
 *     method of TemplateGenerator.
 *
 * Parameters:
 * None
 *
 * Returns:
 * A wrapped error if any is raised by the underlying function it wraps.
 */
func (tg *TemplateGenerator) Scaffold() error {
	// Check if the template already exists?
	_, err := os.Stat(tg.TemplatePath)

	// If the template does not already exist locally, then download it
	if os.IsNotExist(err) {
		fmt.Printf(
			"Template path does not exist...downloading it at: %s\n",
			tg.TemplatePath,
		)

		// TODO: Use the returned filepath returned to extract its contents
		_, err := DownloadTemplate(tg.Owner, tg.Repo)

		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return fmt.Errorf("Error checking template path: %w", err)
	}

	// If the template is found locally, scaffold a project from it
	fmt.Printf("Template found locally at: %s\n", tg.TemplatePath)

	return nil
}
