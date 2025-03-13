// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package m

type CodeModel struct {
	Configurations []Configuration `json:"configurations"`
	Kind           string          `json:"kind"`
	Paths          Paths           `json:"paths"`
	Version        Version         `json:"version"`
}

type Configuration struct {
	Name        string      `json:"name"`
	Directories []Directory `json:"directories"`
	Projects    []Project   `json:"projects"`
	Targets     []Target    `json:"targets"`
}

type Paths struct {
	Build  string `json:"build"`
	Source string `json:"source"`
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
}

type Directory struct {
	Source              string       `json:"source"`
	Build               string       `json:"build"`
	ParentIndex         *int         `json:"parentIndex"`
	ChildIndexes        []int        `json:"childIndexes"`
	ProjectIndex        int          `json:"projectIndex"`
	TargetIndexes       []int        `json:"targetIndexes"`
	MinimumCMakeVersion CMakeVersion `json:"minimumCMakeVersion"`
	HasInstallRule      bool         `json:"hasInstallRule"`
	JSONFile            string       `json:"jsonFile"`
}

type Project struct {
	Name             string `json:"name"`
	ParentIndex      *int   `json:"parentIndex"`
	ChildIndexes     []int  `json:"childIndexes"`
	DirectoryIndexes []int  `json:"directoryIndexes"`
	TargetIndexes    []int  `json:"targetIndexes"`
}

type Target struct {
	Name           string `json:"name"`
	ID             string `json:"id"`
	DirectoryIndex int    `json:"directoryIndex"`
	ProjectIndex   int    `json:"projectIndex"`
	JSONFile       string `json:"jsonFile"`
}
