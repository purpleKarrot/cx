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

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	RunE:  RunInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().String("component", "", "Component-based install. Only install <comp>")

	installCmd.Flags().Bool("strip", false, "Performing install/strip")
	installCmd.Flags().BoolP("verbose", "G", false, "Enable verbose output")

	viper.BindPFlag("install.strip", installCmd.Flags().Lookup("strip"))
	viper.BindPFlag("install.verbose", installCmd.Flags().Lookup("verbose"))
}

func RunInstall(cmd *cobra.Command, args []string) error {
	if err := RunBuild(cmd, args); err != nil {
		return err
	}

	cm := exec.Command("cmake", "--install", filepath.Join(rootBinaryDir, projectSubdir))

	if config := viper.GetString("config"); config != "" {
		cm.Args = append(cm.Args, "--config", config)
	}

	if component := cmd.Flag("component").Value.String(); component != "" {
		cm.Args = append(cm.Args, "--component", component)
	}

	if strip := viper.GetBool("install.strip"); strip {
		cm.Args = append(cm.Args, "--strip")
	}

	if verbose := viper.GetBool("install.verbose"); verbose {
		cm.Args = append(cm.Args, "--verbose")
	}

	cm.Args = append(cm.Args, "--parallel", viper.GetString("parallel"))

	return x.Run(cm, verbose)
}
