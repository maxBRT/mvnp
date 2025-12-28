package generate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func FindRoot(startDir string) (string, error) {
	current := startDir
	for {
		// 1. Check for pom.xml in the current directory
		if _, err := os.Stat(filepath.Join(current, "pom.xml")); err == nil {
			return current, nil
		}

		// 2. Move up one level
		parent := filepath.Dir(current)

		// 3. Stop if we hit the top (parent is same as current)
		if parent == current {
			return "", fmt.Errorf("pom.xml not found")
		}
		current = parent
	}

}

// ListAllPackages returns a slice of all Java packages in the project
// e.g. ["com.myapp.auth", "com.myapp.users", ...]
func ListAllPackages(srcRoot string) ([]string, error) {
	var packages []string

	err := filepath.WalkDir(srcRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 1. Skip non-directories
		if !d.IsDir() {
			return nil
		}

		// 2. Calculate the package name from the path
		relPath, _ := filepath.Rel(srcRoot, path)
		if relPath == "." {
			return nil
		}

		if containsJavaFiles(path) {
			pkgName := strings.ReplaceAll(relPath, string(os.PathSeparator), ".")
			packages = append(packages, pkgName)
		}

		return nil
	})

	return packages, err
}

func containsJavaFiles(dir string) bool {
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".java") {
			return true
		}
	}
	return false
}
