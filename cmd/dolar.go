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

// dolarCmd represents the dolar command
var dolarCmd = &cobra.Command{
	Use:   "dolar",
	Short: "Permite obtener el tipo de cambio del dólar de los Estados Unidos de América",
	Long: `Permite obtener el tipo de cambio del dólar de los Estados Unidos de América.

También permite obtener el histórico del tipo de cambio diario utilizando los parámetros
desde y hasta.

Restricciones:

  - Para obtener el histórico se requiere de los parámetros d y t (desde y hasta), 
    que corresponden a las fechas en formato año-mes-día (2019-01-01).`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		desde, _ := cmd.Flags().GetString("desde")
		hasta, _ := cmd.Flags().GetString("hasta")
		c := api.NewClient(&http.Client{})
		data, resp, err := c.Indicadores.TipoDeCambio("dolar", desde, hasta)
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
	tcCmd.AddCommand(dolarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dolarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dolarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dolarCmd.Flags().StringP("desde", "d", "", "Histórico")
	dolarCmd.Flags().StringP("hasta", "t", "", "Histórico")
	dolarCmd.Flags().BoolP("verbose", "v", false, "Verbose command")
}
