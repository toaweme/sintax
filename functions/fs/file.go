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

// File builds the `file` modifier, which reads a file's contents as a string.
// Unlike the plain global modifiers, `file` is a security boundary and is
// constructed with a caller-supplied allowlist of safe directories, so a caller
// wires it into the engine to control exactly which directories a template may
// read from. Pass an empty allowlist to keep file reads disabled entirely.
//
// The piped value is treated as a path and resolved against each safe dir in
// order. The first dir that contains a readable file wins. A path may only point
// inside a safe dir: anything that escapes via ".." is dropped, so a template
// author cannot read siblings or parents of a safe dir. Absolute paths are not
// an escape hatch either. They are joined onto the safe dir (so "/etc/passwd"
// against safe dir "tpl" resolves to "tpl/etc/passwd"), never to the real root.
// When no safe dir yields the file the modifier returns an os.ErrNotExist error,
// which deliberately does not reveal whether the file existed outside the
// allowlist.
//
// value: string (a path, always resolved relative to a safe dir even if absolute)
// returns: string (the file contents)
//
// example: load a template file from an allowlisted dir
// in:  safeDirs = ["./templates"]
// tpl: {{ "greeting.tpl.mdx" | file }}
// out: <contents of ./templates/greeting.tpl.mdx>
//
// example: read a nested path inside the allowlist
// in:  safeDirs = ["./templates"]
// tpl: {{ "emails/welcome.txt" | file }}
// out: <contents of ./templates/emails/welcome.txt>
//
// example: a traversal attempt is rejected, not served
// in:  safeDirs = ["./templates"]
// tpl: {{ "../secrets.env" | file }}
// out: <error: file "../secrets.env" not found>
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
