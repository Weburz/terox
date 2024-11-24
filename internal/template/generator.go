package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

type TemplateGenerator struct {
	Repo         string
	Owner        string
	TemplatePath string
}

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
		Repo:         parts[0],
		Owner:        parts[1],
		TemplatePath: templatePath,
	}, nil
}

func (tg *TemplateGenerator) Scaffold() error {
	// Check if the template already exists?
	_, err := os.Stat(tg.TemplatePath)

	// If the template does not already exist locally, then download it
	if os.IsNotExist(err) {
		fmt.Printf(
			"Template path does not exist...downloading it at: %s\n",
			tg.TemplatePath,
		)
		// TODO: Implementing the downloading logic
		return nil
	} else if err != nil {
		return fmt.Errorf("Error checking template path: %w", err)
	}

	// If the template is found locally, scaffold a project from it
	fmt.Printf("Template found locally at: %s\n", tg.TemplatePath)

	return nil
}
