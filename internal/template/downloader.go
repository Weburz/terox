/**
 * Package template - The "template" package contains the logic to
 * download/scaffold the project from a pre-defined template.
 *
 * This file in particular which is part of the "template" package is
 * responsible for downloading the template from GitHub.
 *
 * TODO: A future update will allow the downloading logic to also fetch
 * templates from locations other than GitHub.
 */
package template

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
 * DownloadTemplate - Download a template from a GitHub repository.
 *
 * Parameters:
 * owner string: The user/organization who owns the template (to be downloaded).
 * repo string: The template's repository (on GitHub).
 *
 * Returns:
 * Returns the filepath to the temporary file where the zipball is downloaded
 * to or raise an error if there exists one.
 */
func DownloadTemplate(owner, repo string) (string, error) {
	// Create an URL string to fetch the HTTP response from
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/zipball",
		owner,
		repo,
	)

	// Get the HTTP response from the URL
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("Failed to download the template: %w", err)
	}

	// Close the HTTP connection after the response is written to disk
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad server response: %d", resp.StatusCode)
	}

	// Create an object to write the HTTP response to
	tempFile, err := os.CreateTemp("", "repoforge-*.zip")

	if err != nil {
		return "", fmt.Errorf("Failed to create the template zipball: %w", err)
	}

	// Close the object after its written to disk
	defer tempFile.Close()

	// Copy the objects stored in the "file" object to disk (or throw an error)
	if _, err := io.Copy(tempFile, resp.Body); err != nil {
		return "", fmt.Errorf("Failed to write the zipball: %w", err)
	}

	// Return the filepath to the temporary zip archive if no errors were raised
	return tempFile.Name(), nil
}
