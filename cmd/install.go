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

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Run:   RunInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().String("config", "", "For multi-configuration tools, choose <cfg>")
	installCmd.Flags().String("component", "", "Component-based install. Only install <comp>")

	installCmd.Flags().Bool("strip", false, "Performing install/strip")
	installCmd.Flags().BoolP("verbose", "G", false, "Enable verbose output")

	viper.BindPFlag("install.strip", installCmd.Flags().Lookup("strip"))
	viper.BindPFlag("install.verbose", installCmd.Flags().Lookup("verbose"))
}

func RunInstall(cmd *cobra.Command, args []string) {
	RunBuild(cmd, args)

	cm := exec.Command("cmake", "--install", filepath.Join(rootBinaryDir, projectSubdir))

	cm.Stdout = os.Stdout
	cm.Stderr = os.Stderr

	if config := cmd.Flag("config").Value.String(); config != "" {
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

	fmt.Printf("\nExecuting command: %s %s\n\n", cm.Path, strings.Join(cm.Args[1:], " "))
	if err := cm.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}
}
