// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"os/exec"

	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the generated project in the associated application",
	Run: func(cmd *cobra.Command, args []string) {
		RequireConfigure(cmd, args)
		x.Run(exec.Command("cmake", "--open", rootBinaryDir), verbose)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
