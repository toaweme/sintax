package sintax

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type Template struct {
	Data any
}

func Files(cwd string) (files []string, err error) {
	filepath.Walk(cwd, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking path %s: %w", path, err)
		}

		files = append(files, path)
		return nil
	})

	return files, nil
}

func Dirs(cwd string) (dirs []string, err error) {
	filepath.Walk(cwd, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking path %s: %w", path, err)
		}

		if info.IsDir() {
			dirs = append(dirs, path)
		}

		return nil
	})

	return dirs, nil
}

// search relative structures
