package sitemap

import (
	"fmt"
	"os"
	"path/filepath"
)

func ensureOutputDir(relativePath string) (string, error) {
	if relativePath == "" {
		return "", fmt.Errorf("sitemap: output directory path is empty")
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("sitemap: get working directory: %w", err)
	}

	dir := filepath.Join(cwd, relativePath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("sitemap: create directory %q: %w", dir, err)
	}

	return dir, nil
}
