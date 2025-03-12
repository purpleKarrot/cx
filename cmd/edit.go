// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a cache editor",
	Run:   RunEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().Bool("gui", false, "Use Qt GUI")
	viper.BindPFlag("edit.gui", editCmd.Flags().Lookup("gui"))
}

func RunEdit(cmd *cobra.Command, args []string) {
	cm := MakeConfigureCmd(x.If(viper.GetBool("edit.gui"), "cmake-gui", "ccmake"))

	cm.Stdin = os.Stdin
	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr

	fmt.Printf("\nExecuting command: %s %s\n\n", cm.Path, strings.Join(cm.Args[1:], " "))
	if err := cm.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
}
