// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete the existing build directory",
	RunE:  RunClean,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func RunClean(cmd *cobra.Command, args []string) error {
	if verbose {
		fmt.Printf("Cleaning build directory %s\n", rootBinaryDir)
	}

	return os.RemoveAll(rootBinaryDir)
}
