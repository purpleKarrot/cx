// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose       bool
	rootSourceDir string
	rootBinaryDir string
	projectSubdir string
)

var rootCmd = &cobra.Command{
	Use:   "cx",
	Short: "A brief description of your application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initProjectRoot, initConfig)

	flags := rootCmd.PersistentFlags()
	flags.BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

	flags.String("config", "", "Specify the build configuration")
	viper.BindPFlag("config", flags.Lookup("config"))
}

func initProjectRoot() {
	startDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	rootSourceDir, err = findProjectRoot(startDir)
	if err != nil {
		log.Fatalf("Failed to find project root: %v", err)
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(rootSourceDir)))
	rootBinaryDir = filepath.Join(xdg.CacheHome, "cx", hash)

	projectSubdir, err = filepath.Rel(rootSourceDir, startDir)
	if err != nil {
		log.Fatalf("Failed to calculate subdir: %v", err)
	}
}

func initConfig() {
	viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "cx"))
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func findProjectRoot(startDir string) (string, error) {
	var rootDir string
	dir := startDir
	for {
		if _, err := os.Stat(filepath.Join(dir, "CMakeLists.txt")); err == nil {
			rootDir = dir
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	if rootDir == "" {
		return "", fmt.Errorf("CMakeLists.txt not found")
	}
	return rootDir, nil
}
