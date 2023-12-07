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

// pescaCmd represents the pesca command
var pescaCmd = &cobra.Command{
	Use:   "pesca",
	Short: "Permite consultar información relacionada al productor pesquero",
	Long: `Permite obtener los siguientes datos:

  - Del registro de productores agropecuarios del MAG: el nombre, el estado, el indicador de si está activo y la fecha de baja. 
  - Del registro de INCOPESCA: el indicador de si está activo, el nombre del permisionario y la fecha de vencimiento. 
  - Del registro de acuicultores: el indicador de si está activo, el nombre del acuicultor y la fecha de vencimiento.

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
		data, resp, err := c.FacturaElectronica.Pesca(identification)
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
	feCmd.AddCommand(pescaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pescaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pescaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pescaCmd.Flags().StringP("id", "i", "0", "Número de identificación del productor pesquero")
	pescaCmd.Flags().BoolP("verbose", "v", false, "Verbose command")
}
