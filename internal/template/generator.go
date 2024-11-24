package template

import (
	"fmt"
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
	// TODO: Add the scaffolding logic here
	return nil
}
