// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os"

	"github.com/purpleKarrot/cx/m"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "clean",
		Short: "Delete the existing build directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths, err := m.FindProjectPaths()
			if err != nil {
				return err
			}

			if verbose {
				fmt.Printf("Cleaning build directory %s\n", paths.Binary)
			}

			return os.RemoveAll(paths.Binary)
		},
	})
}
