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

// agropecuarioCmd represents the agropecuario command
var agropecuarioCmd = &cobra.Command{
	Use:     "agropecuario",
	Aliases: []string{"agro"},
	Short:   "Permite consultar información relacionada al productor agropecuario",
	Long: `Permite obtener el nombre, el estado, el indicador de si está activo
y la fecha de baja de los productores agropecuarios en el MAG, usando como 
parámetro el número de "identificacion" (sin hacer uso de guiones). 

Para mayor practicidad, se incluye también el resultado de la situación tributaria del comando de actividad económica (hacienda fe ae).

Identificaciones soportadas:

  1. Físicas nacionales
  2. Jurídicas nacionales
  3. DIMEX

Restricciones:

  - Físicas: debe contener 9 números, sin cero al inicio y sin guiones.
  - Jurídicas: debe contener 10 dígitos, sin guiones y sin utilizar los dos últimos dígitos verificadores.
  - DIMEX: debe contener 11 o 12 dígitos, sin cero al inicio y sin guiones.
`,
	Run: func(cmd *cobra.Command, args []string) {
		identification, _ := cmd.Flags().GetString("id")
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := api.NewClient(&http.Client{})
		data, resp, err := c.FacturaElectronica.Agropecuario(identification)
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
	feCmd.AddCommand(agropecuarioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agropecuarioCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agropecuarioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	agropecuarioCmd.Flags().StringP("id", "i", "0", "Número de identificación del productor agropecuario")
	agropecuarioCmd.Flags().BoolP("verbose", "v", false, "Verbose command")
}
