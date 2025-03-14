// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"os/exec"

	"github.com/purpleKarrot/cx/m"
	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "open",
		Short: "Open the generated project in the associated application",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths, err := m.FindProjectPaths()
			if err != nil {
				return err
			}

			if err := RequireConfigure(cmd, args); err != nil {
				return err
			}

			return x.Run(exec.Command("cmake", "--open", paths.Binary), verbose)
		},
	})
}
