// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the generated project in the associated application",
	Run:   RunOpen,
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func RunOpen(cmd *cobra.Command, args []string) {
	RequireConfigure(cmd, args)

	cm := exec.Command("cmake", "--open", rootBinaryDir)
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr

	if err := cm.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}
}
