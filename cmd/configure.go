// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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

func RunConfigure(cmd *cobra.Command, args []string) {
	cm := exec.Command("cmake", "-S"+rootSourceDir, "-B"+rootBinaryDir)

	generator := viper.GetString("generator")
	if generator != "" {
		cm.Args = append(cm.Args, "-G"+generator)
	}

	if toolset := viper.GetString("toolset"); toolset != "" {
		cm.Args = append(cm.Args, "-T"+toolset)
	}

	if platform := viper.GetString("platform"); platform != "" {
		cm.Args = append(cm.Args, "-A"+platform)
	}

	// TODO: Don't specify build type for multi-config generators
	if buildType := viper.GetString("build_type"); buildType != "" {
		cm.Args = append(cm.Args, "-DCMAKE_BUILD_TYPE="+buildType)
	}

	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr

	if err := cm.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	os.Mkdir(filepath.Join(rootBinaryDir, ".cx"), 0755)
	viper.WriteConfigAs(filepath.Join(rootBinaryDir, ".cx", "config.json"))
}

func RequireConfigure(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(filepath.Join(rootBinaryDir, "CMakeCache.txt")); os.IsNotExist(err) {
		RunConfigure(cmd, args)
	}
}
