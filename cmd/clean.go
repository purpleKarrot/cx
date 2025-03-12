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
	Run:   RunClean,
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func RunClean(cmd *cobra.Command, args []string) {
	if err := os.RemoveAll(rootBinaryDir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to delete directory: %v\n", err)
		os.Exit(1)
	}
}
