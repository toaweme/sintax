// Package fs provides a modifier that reads file contents from an allowlist.
package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFile is the template name for the File modifier.
const ModifierNameFile functions.ModifierName = "file"

// File builds the `file` modifier, which reads a file's contents as a string so
// a template can embed on-disk content (for example piping the result into the
// template modifier as a partial). Unlike the plain global modifiers, `file` is
// a security boundary. The application that wires it in supplies an allowlist of
// safe directories that controls exactly which directories a template may read
// from, and an empty allowlist disables file reads entirely.
//
// The path is resolved against each safe dir in order and the first dir that
// contains a readable file wins. A path may only point inside a safe dir.
// Anything that escapes via ".." is dropped, so a template author cannot read
// siblings or parents of a safe dir. Absolute paths are not an escape hatch
// either. They are joined onto the safe dir (so "/etc/passwd" against safe dir
// "tpl" resolves to "tpl/etc/passwd"), never to the real root. When no safe dir
// yields the file it returns a not-found error that deliberately does not reveal
// whether the file existed outside the allowlist.
func File(safeDirs []string) func(path string) (string, error) {
	return func(path string) (string, error) {
		paths, err := resolveSafePaths(path, safeDirs)
		if err != nil {
			return "", err
		}

		for _, full := range paths {
			data, err := os.ReadFile(full)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return "", fmt.Errorf("failed to read file %q: %w", path, err)
			}
			return string(data), nil
		}

		return "", fmt.Errorf("failed to read file %q: %w", path, os.ErrNotExist)
	}
}

// resolveSafePaths validates path against the configured safe dirs, returning
// the cleaned candidate paths to read (one per safe dir the path stays inside).
// It performs no I/O. Paths that escape their safe dir via ".." are dropped, and
// if none remain it returns os.ErrNotExist so a traversal attempt is
// indistinguishable from a genuine miss.
func resolveSafePaths(path string, safeDirs []string) (paths []string, err error) {
	if len(safeDirs) == 0 {
		return nil, fmt.Errorf("failed to read file %q: no safe directories configured", path)
	}

	for _, dir := range safeDirs {
		cleanDir := filepath.Clean(dir)
		full := filepath.Clean(filepath.Join(cleanDir, path))

		// reject anything that escapes the safe dir via ".."
		if full != cleanDir && !strings.HasPrefix(full, cleanDir+string(os.PathSeparator)) {
			continue
		}
		paths = append(paths, full)
	}

	if len(paths) == 0 {
		return nil, fmt.Errorf("failed to read file %q: %w", path, os.ErrNotExist)
	}
	return paths, nil
}
