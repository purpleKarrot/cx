// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build all targets in the current directory",
	Run:   RunBuild,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func RunBuild(cmd *cobra.Command, args []string) {
	RequireConfigure(cmd, args)

	var cm *exec.Cmd
	generator := viper.GetString("generator")
	if strings.Contains(generator, "Ninja") {
		all := x.If(projectSubdir != ".", projectSubdir+"/all", "all")
		cm = exec.Command("cmake", "--build", rootBinaryDir, "--target", all)
	} else if generator == "Xcode" {
		cm = exec.Command("cmake", "--build", rootBinaryDir, "--target", "ALL_BUILD")
	} else {
		cm = exec.Command("cmake", "--build", filepath.Join(rootBinaryDir, projectSubdir))
	}

	// TODO: Don't specify build type for single config generators
	if config := viper.GetString("config"); config != "" {
		cm.Args = append(cm.Args, "--config", config)
	}

	x.Run(cm, verbose)
}
