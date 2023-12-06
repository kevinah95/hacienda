/*
 * Copyright (c) 2023 Kevin Hernández Rostrán
 * Use of this source code is governed by the Apache 2.0 license that can be found in the LICENSE file.
 */

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hacienda",
	Short: "Hacienda es una aplicación que se conecta al API del Ministerio de Hacienda de Costa Rica.",
	Long: `Hacienda expone la interfaz de línea de comandos para el uso del API del
Ministerio de Hacienda de Costa Rica (https://www.hacienda.go.cr).

Más información sobre el API del Ministerio de Hacienda en: 
	
	https://api.hacienda.go.cr/status o https://bit.ly/2McNcpX
`,
	Version: "0.0.1", // x-release-please-version
	Example: `
  hacienda fe ae --id=2100042005
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hacienda.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
