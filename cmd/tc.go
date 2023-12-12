/*
 * Copyright (c) 2023 Kevin Hernández Rostrán
 * Use of this source code is governed by the Apache 2.0 license that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"github.com/kevinah95/hacienda/api"
	"net/http"

	"github.com/spf13/cobra"
)

// tcCmd represents the tc command
var tcCmd = &cobra.Command{
	Use:     "tc",
	Aliases: []string{"tipocambio"},
	Short:   "Permite obtener el tipo de cambio del dólar y del Euro",
	Long: `Permite obtener el tipo de cambio del dólar de los Estados Unidos de América
y del Euro (dólares y colones).  No requiere de ningún tipo de parámetro.`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := api.NewClient(&http.Client{})
		data, resp, err := c.Indicadores.TipoDeCambio("", "", "")
		if err != nil {
			return
		}

		if verbose {
			fmt.Println(resp)
		}
		fmt.Println(data)
	},
}

func init() {
	rootCmd.AddCommand(tcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tcCmd.Flags().BoolP("verbose", "v", false, "Verbose command")
}
