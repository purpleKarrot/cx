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

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Generate a build system",
	Run:   RunConfigure,
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringP("generator", "G", "", "Specify a build system generator")
	configureCmd.Flags().StringP("toolset", "T", "", "Specify toolset name if supported by generator")
	configureCmd.Flags().StringP("platform", "A", "", "Specify platform name if supported by generator")

	viper.BindPFlag("generator", configureCmd.Flags().Lookup("generator"))
	viper.BindPFlag("toolset", configureCmd.Flags().Lookup("toolset"))
	viper.BindPFlag("platform", configureCmd.Flags().Lookup("platform"))
}

func MakeConfigureCmd(cmake string) *exec.Cmd {
	if _, err := os.Stat(filepath.Join(rootBinaryDir, "CMakeCache.txt")); !os.IsNotExist(err) {
		return exec.Command(cmake, rootBinaryDir)
	}

	cmd := exec.Command(cmake, "-S"+rootSourceDir, "-B"+rootBinaryDir)

	generator := viper.GetString("generator")
	if generator != "" {
		cmd.Args = append(cmd.Args, "-G"+generator)
	}

	if toolset := viper.GetString("toolset"); toolset != "" {
		cmd.Args = append(cmd.Args, "-T"+toolset)
	}

	if platform := viper.GetString("platform"); platform != "" {
		cmd.Args = append(cmd.Args, "-A"+platform)
	}

	// TODO: use `--config` once it is supported by CMake.
	// See: https://gitlab.kitware.com/cmake/cmake/-/merge_requests/10387
	// The advantage is that it will be ignored for multi-config generators.
	if config := viper.GetString("config"); config != "" {
		cmd.Args = append(cmd.Args, "-DCMAKE_BUILD_TYPE="+config)
	}

	return cmd
}

func RunConfigure(cmd *cobra.Command, args []string) {
	cm := MakeConfigureCmd("cmake")

	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr

	fmt.Printf("\nExecuting command: %s %s\n\n", cm.Path, strings.Join(cm.Args[1:], " "))
	if err := cm.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
}

func RequireConfigure(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(filepath.Join(rootBinaryDir, "CMakeCache.txt")); os.IsNotExist(err) {
		RunConfigure(cmd, args)
	}
}
