// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package cmd

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verbose bool

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
	cobra.OnInitialize(initConfig)

	flags := rootCmd.PersistentFlags()
	flags.BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

	flags.String("config", "Debug", "Specify the build configuration")
	viper.BindPFlag("config", flags.Lookup("config"))

	flags.Bool("fresh", false, "Configure a fresh build tree, removing any existing cache file.")

	flags.IntP("parallel", "j", runtime.NumCPU(), "Specify the number of parallel jobs")
	viper.BindPFlag("parallel", flags.Lookup("parallel"))
}

func initConfig() {
	viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "cx"))
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
