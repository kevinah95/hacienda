/*
 * Copyright (c) 2023 Kevin Hernández Rostrán
 * Use of this source code is governed by the Apache 2.0 license that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"github.com/kevinah95/hacienda/api"
	"github.com/spf13/cobra"
	"net/http"
)

// aeCmd represents the ae command
var aeCmd = &cobra.Command{
	Use:   "ae",
	Short: "Permite consultar información relacionada al contribuyente",
	Long: `Permite obtener el nombre, el tipo de identificación, el régimen, 
la situación tributaria y las actividades económicas asociadas a un contribuyente, 
usando como parámetro el número de identificación (sin hacer uso de guiones).

Identificaciones soportadas:

  1. Físicas nacionales
  2. Jurídicas nacionales
  3. DIMEX: Documento de Identificación Migratorio para Extranjeros
  4. NITE: Numero de Identificación Tributario Especial

Restricciones:

  - Físicas: debe contener 9 números, sin cero al inicio y sin guiones.
  - DIMEX: debe contener 11 o 12 dígitos, sin cero al inicio y sin guiones.
  - Jurídicas y NITE: debe contener 10 dígitos, sin guiones y sin utilizar los dos últimos dígitos verificadores.

Ver también:

  https://atv.hacienda.go.cr/ATV/frmConsultaSituTributaria.aspx
`,
	Run: func(cmd *cobra.Command, args []string) {
		identification, _ := cmd.Flags().GetString("id")
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := api.NewClient(&http.Client{})
		data, resp, err := c.FacturaElectronica.ActividadEconomica(identification)
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
	feCmd.AddCommand(aeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	aeCmd.Flags().StringP("id", "i", "0", "Número de identificación del contribuyente")
	aeCmd.Flags().BoolP("verbose", "v", false, "Verbose command")
}
