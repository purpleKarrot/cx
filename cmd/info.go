// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"encoding/json"
	"os"

	"github.com/purpleKarrot/cx/m"
	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
)

func init() {
	type Info struct {
		CMake *m.CMake `json:"cmake,omitempty"`
		Paths *m.Paths `json:"paths"`
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Print information about the generated project",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := m.LoadIndex(rootBinaryDir)
			return json.NewEncoder(os.Stdout).Encode(&Info{
				CMake: x.If(err == nil, &api.CMake, nil),
				Paths: &m.Paths{
					Source: rootSourceDir,
					Build:  rootBinaryDir,
				},
			})
		},
	})
}
