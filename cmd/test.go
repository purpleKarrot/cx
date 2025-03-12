// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"os/exec"
	"path/filepath"

	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Run:   RunTest,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func RunTest(cmd *cobra.Command, args []string) {
	RunBuild(cmd, args)

	cm := exec.Command("ctest", "--test-dir", filepath.Join(rootBinaryDir, projectSubdir), "--parallel")

	// TODO: Don't specify build type for single config generators
	if buildType := viper.GetString("build_type"); buildType != "" {
		cm.Args = append(cm.Args, "-C", buildType)
	}

	x.Run(cm)
}
