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

// euroCmd represents the euro command
var euroCmd = &cobra.Command{
	Use:   "euro",
	Short: "Permite obtener el tipo de cambio del Euro (dólares y colones)",
	Long: `Permite obtener el tipo de cambio del Euro (dólares y colones).

No requiere de ningún tipo de parámetro.`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := api.NewClient(&http.Client{})
		data, resp, err := c.Indicadores.TipoDeCambio("euro", "", "")
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
	tcCmd.AddCommand(euroCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// euroCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// euroCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	euroCmd.Flags().BoolP("verbose", "v", false, "Verbose command")
}
