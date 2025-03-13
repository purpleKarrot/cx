// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a cache editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmake := x.If(viper.GetBool("edit.gui"), "cmake-gui", "ccmake")
		c, err := MakeConfigureCmd(cmake)
		if err != nil {
			return err
		}
		return x.Run(c, verbose)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().Bool("gui", false, "Use Qt GUI")
	viper.BindPFlag("edit.gui", editCmd.Flags().Lookup("gui"))
}
