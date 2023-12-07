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

var verbose bool
var Identification string

// exCmd represents the ex command
var exCmd = &cobra.Command{
	Use:   "ex",
	Short: "Permite consultar la información correspondiente a una exoneración",
	Long: `Permite obtener la información correspondiente a una exoneración.
Requiere el parámetro "autorizacion" cuyo formato debe seguir la regla "al-00000000-00".

Cuando una autorización posee CABYS asociados, el campo "poseeCabys" tendrá un valor "true" y en consecuencia 
aparecerá el campo "cabys" que corresponde a un array de códigos CABYS. 
Si el valor de "poseeCabys" es "false", la respuesta no incluye el arreglo de códigos.

Restricciones:

  - Si el documento de exoneración no existe se devuelve el código "404".  
  - Si el formato del parámetro "autorizacion" es incorrecto o no existe como parámetro, 
    se devuelve el código "400".


`,
	Run: func(cmd *cobra.Command, args []string) {
		authz, _ := cmd.Flags().GetString("authz")
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := api.NewClient(&http.Client{})
		data, resp, err := c.FacturaElectronica.Exoneracion(authz)
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
	feCmd.AddCommand(exCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	exCmd.Flags().StringP("authz", "a", "al-00000000-00", "Número de documento.")
	exCmd.Flags().BoolP("verbose", "v", false, "verbose command")
}
