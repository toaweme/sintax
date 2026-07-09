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

// File builds a modifier that reads a file's contents as a string. The path is
// resolved against safeDirs (in order) and may only point inside one of them:
// any path that escapes a safe dir via ".." is rejected. This is the closure a
// caller wires into the engine to control which directories templates can read.
//
// value: string (a path, relative to a safe dir or absolute inside one)
// returns: string (the file contents)
//
// example: load a template file from an allowlisted dir
// in:  safeDirs = ["./templates"]
// tpl: {{ "greeting.tpl.mdx" | file }}
// out: <contents of ./templates/greeting.tpl.mdx>
func File(safeDirs []string) func(value any, params []any) (any, error) {
	return func(value any, params []any) (any, error) {
		rel, paths, err := resolveSafePaths(value, safeDirs)
		if err != nil {
			return nil, err
		}

		for _, full := range paths {
			data, err := os.ReadFile(full)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return nil, fmt.Errorf("failed to read file %q: %w", rel, err)
			}
			return string(data), nil
		}

		return nil, fmt.Errorf("failed to read file %q: %w", rel, os.ErrNotExist)
	}
}

// resolveSafePaths validates the file argument and the configured safe dirs,
// returning the requested path and the cleaned candidate paths to read (one per
// safe dir the path stays inside). It performs no I/O. Paths that escape their
// safe dir via ".." are dropped; if none remain it returns os.ErrNotExist.
func resolveSafePaths(value any, safeDirs []string) (rel string, paths []string, err error) {
	rel, err = functions.ValueString(value)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read file path: %w", err)
	}
	if len(safeDirs) == 0 {
		return rel, nil, fmt.Errorf("failed to read file %q: no safe directories configured", rel)
	}

	for _, dir := range safeDirs {
		cleanDir := filepath.Clean(dir)
		full := filepath.Clean(filepath.Join(cleanDir, rel))

		// reject anything that escapes the safe dir via ".."
		if full != cleanDir && !strings.HasPrefix(full, cleanDir+string(os.PathSeparator)) {
			continue
		}
		paths = append(paths, full)
	}

	if len(paths) == 0 {
		return rel, nil, fmt.Errorf("failed to read file %q: %w", rel, os.ErrNotExist)
	}
	return rel, paths, nil
}
