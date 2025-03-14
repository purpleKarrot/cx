// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package m

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

type ProjectPaths struct {
	Source string
	Binary string
	Subdir string
}

func FindProjectPaths() (*ProjectPaths, error) {
	startDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Failed to get current working directory: %v", err)
	}

	source, err := findProjectRoot(startDir)
	if err != nil {
		return nil, fmt.Errorf("Failed to find project root: %v", err)
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(source)))
	binary := filepath.Join(xdg.CacheHome, "cx", hash)

	subdir, err := filepath.Rel(source, startDir)
	if err != nil {
		return nil, fmt.Errorf("Failed to calculate subdir: %v", err)
	}

	return &ProjectPaths{source, binary, subdir}, nil
}

func findProjectRoot(startDir string) (string, error) {
	var rootDir string
	dir := startDir
	for {
		if _, err := os.Stat(filepath.Join(dir, "CMakeLists.txt")); err == nil {
			rootDir = dir
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	if rootDir == "" {
		return "", fmt.Errorf("CMakeLists.txt not found")
	}
	return rootDir, nil
}
