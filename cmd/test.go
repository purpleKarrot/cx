// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"os/exec"
	"path/filepath"

	"github.com/purpleKarrot/cx/m"
	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	RunE:  RunTest,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func RunTest(cmd *cobra.Command, args []string) error {
	api, err := m.LoadIndex(rootBinaryDir)
	if err != nil {
		return err
	}

	if err := RunBuild(cmd, args); err != nil {
		return err
	}

	cm := exec.Command("ctest", "--test-dir", filepath.Join(rootBinaryDir, projectSubdir))

	if api.CMake.Generator.MultiConfig {
		cm.Args = append(cm.Args, "-C", viper.GetString("config"))
	}

	cm.Args = append(cm.Args, "--parallel", viper.GetString("parallel"))

	return x.Run(cm, verbose)
}
