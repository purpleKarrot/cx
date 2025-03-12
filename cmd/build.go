// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
	if generator == "Ninja" {
		cm = exec.Command("cmake", "--build", rootBinaryDir, "--target", ninjaAllTarget())
	} else if generator == "Xcode" {
		cm = exec.Command("cmake", "--build", rootBinaryDir, "--target", "ALL_BUILD")
	} else {
		cm = exec.Command("cmake", "--build", filepath.Join(rootBinaryDir, projectSubdir))
	}

	// TODO: Don't specify build type for single config generators
	if buildType := viper.GetString("build_type"); buildType != "" {
		cm.Args = append(cm.Args, "--config", buildType)
	}

	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr

	fmt.Printf("\nExecuting command: %s %s\n\n", cm.Path, strings.Join(cm.Args[1:], " "))
	if err := cm.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}
}

func ninjaAllTarget() string {
	if projectSubdir != "." {
		return projectSubdir + "/all"
	}
	return "all"
}
