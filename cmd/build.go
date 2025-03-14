// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/purpleKarrot/cx/m"
	"github.com/purpleKarrot/cx/x"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build all targets in the current directory",
	RunE:  RunBuild,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func RunBuild(cmd *cobra.Command, args []string) error {
	paths, err := m.FindProjectPaths()
	if err != nil {
		return err
	}

	if err := RequireConfigure(cmd, args); err != nil {
		return err
	}

	api, err := m.LoadIndex(paths.Binary)
	if err != nil {
		return err
	}

	model, err := api.LoadCodeModel(paths.Binary)
	if err != nil {
		return err
	}

	dir := model.FindDirectory(viper.GetString("config"), paths.Subdir)
	if dir == nil {
		fmt.Printf("Directory not found\n")
		return nil
	}

	var cm *exec.Cmd
	generator := viper.GetString("generator")
	if strings.Contains(generator, "Ninja") {
		all := x.If(dir.Build != ".", dir.Build+"/all", "all")
		cm = exec.Command("cmake", "--build", paths.Binary, "--target", all)
	} else if generator == "Xcode" {
		cm = exec.Command("cmake", "--build", paths.Binary, "--target", "ALL_BUILD")
	} else {
		cm = exec.Command("cmake", "--build", filepath.Join(paths.Binary, dir.Build))
	}

	if api.CMake.Generator.MultiConfig {
		cm.Args = append(cm.Args, "--config", viper.GetString("config"))
	}

	cm.Args = append(cm.Args, "--parallel", viper.GetString("parallel"))

	return x.Run(cm, verbose)
}
