// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Generate a build system",
	RunE:  RunConfigure,
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

func MakeConfigureCmd(cmake string) (*exec.Cmd, error) {
	if _, err := os.Stat(filepath.Join(rootBinaryDir, "CMakeCache.txt")); !os.IsNotExist(err) {
		return exec.Command(cmake, rootBinaryDir), nil
	}

	api := filepath.Join(rootBinaryDir, ".cmake", "api", "v1", "query", "client-cx")
	if err := os.MkdirAll(api, 0755); err != nil {
		return nil, fmt.Errorf("Failed to create directory %s: %v", api, err)
	}

	file, _ := os.Create(filepath.Join(api, "codemodel-v2"))
	file.Close()

	cmd := exec.Command(cmake, rootSourceDir)
	cmd.Dir = rootBinaryDir

	generator := viper.GetString("generator")
	if generator != "" {
		cmd.Args = append(cmd.Args, "-G", generator)
	}

	if toolset := viper.GetString("toolset"); toolset != "" {
		cmd.Args = append(cmd.Args, "-T", toolset)
	}

	if platform := viper.GetString("platform"); platform != "" {
		cmd.Args = append(cmd.Args, "-A", platform)
	}

	// TODO: use `--config` once it is supported by CMake.
	// See: https://gitlab.kitware.com/cmake/cmake/-/merge_requests/10387
	// The advantage is that it will be ignored for multi-config generators.
	if config := viper.GetString("config"); config != "" {
		cmd.Args = append(cmd.Args, "-DCMAKE_BUILD_TYPE="+config)
	}

	return cmd, nil
}

func RunConfigure(cmd *cobra.Command, args []string) error {
	c, err := MakeConfigureCmd("cmake")
	if err != nil {
		return err
	}

	if fresh, _ := cmd.Flags().GetBool("fresh"); fresh {
		c.Args = append(c.Args, "--fresh")
	}

	return x.Run(c, verbose)
}

func RequireConfigure(cmd *cobra.Command, args []string) error {
	needed := func() bool {
		if fresh, _ := cmd.Flags().GetBool("fresh"); fresh {
			return true
		}
		if _, err := os.Stat(filepath.Join(rootBinaryDir, "CMakeCache.txt")); os.IsNotExist(err) {
			return true
		}
		return false
	}
	if needed() {
		return RunConfigure(cmd, args)
	}
	return nil
}
