// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package m

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type Index struct {
	CMake CMake          `json:"cmake"`
	Reply map[string]any `json:"reply"`
}

type CMake struct {
	Generator CMakeGenerator `json:"generator"`
	Paths     CMakePaths     `json:"paths"`
	Version   CMakeVersion   `json:"version"`
}

type CMakeGenerator struct {
	Name        string `json:"name"`
	MultiConfig bool   `json:"multiConfig"`
}

type CMakePaths struct {
	CMake string `json:"cmake"`
	CTest string `json:"ctest"`
	CPack string `json:"cpack"`
	Root  string `json:"root"`
}

type CMakeVersion struct {
	IsDirty bool   `json:"isDirty"`
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
	Patch   int    `json:"patch"`
	String  string `json:"string"`
	Suffix  string `json:"suffix"`
}

func LoadIndex(path string) (*Index, error) {
	pattern := filepath.Join(path, ".cmake", "api", "v1", "reply", "index-*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("no matching files found")
	}

	sort.Strings(matches)
	lastMatch := matches[len(matches)-1]

	data, err := os.ReadFile(lastMatch)
	if err != nil {
		return nil, err
	}

	var index Index
	if err = json.Unmarshal(data, &index); err != nil {
		return nil, err
	}

	return &index, nil
}

func (index *Index) LoadCodeModel(path string) (*CodeModel, error) {
	clientCX := index.Reply["client-cx"].(map[string]any)
	codemodelV2 := clientCX["codemodel-v2"].(map[string]any)
	jsonFile := codemodelV2["jsonFile"].(string)
	filePath := filepath.Join(path, ".cmake", "api", "v1", "reply", jsonFile)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var codeModel CodeModel
	if err := json.Unmarshal(data, &codeModel); err != nil {
		return nil, err
	}

	return &codeModel, nil
}
