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

// cabysCmd represents the cabys command
var cabysCmd = &cobra.Command{
	Use: "cabys",
	Example: `  hacienda fe cabys --descripcion="Jugo de tomate"
  hacienda fe cabys --codigo="2132100000100"
  hacienda fe cabys --descripcion="Jugo de tomate" --top=2`,
	Short: "Permite obtener la información correspondiente al Catálogo de Bienes y Servicios (CABYS)",
	Long: `Permite obtener la información correspondiente al Catalago de Bienes y Servicios (CABYS, https://www.hacienda.go.cr/docs/Catalogobienesyservicios.pdf),
a partir de la descripción de los bienes y servicios o su número de código.  

Puede utilizar los parámetros descripcion o codigo de la siguiente manera:

  - Por descripción del bien o servicio:
	
    hacienda fe cabys --descripcion="Jugo de tomate"
  
  - Por el código del bien o servicio:
	
    hacienda fe cabys --codigo="2132100000100"
  
  - El parámetro q puede ser usado en combinación con el parámetro top 
    para hacer una búsqueda limitada de bienes y servicios de la siguiente forma:
	
    hacienda fe cabys --descripcion="Jugo de tomate" --top=2
`,
	Run: func(cmd *cobra.Command, args []string) {
		codigo, _ := cmd.Flags().GetString("codigo")
		descripcion, _ := cmd.Flags().GetString("descripcion")
		top, _ := cmd.Flags().GetString("top")
		verbose, _ := cmd.Flags().GetBool("verbose")

		c := api.NewClient(&http.Client{})
		data, resp, err := c.FacturaElectronica.Cabys(codigo, descripcion, top)
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
	feCmd.AddCommand(cabysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cabysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cabysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cabysCmd.Flags().StringP("descripcion", "d", "", "Descripción del bien o servicio")
	cabysCmd.Flags().StringP("top", "t", "", "Limite de resultados para el parámetro descripción")
	cabysCmd.Flags().StringP("codigo", "c", "", "Código del bien o servicio")
	cabysCmd.Flags().BoolP("verbose", "v", false, "Verbose command")
	cabysCmd.MarkFlagsOneRequired("codigo", "descripcion")
	cabysCmd.MarkFlagsMutuallyExclusive("codigo", "descripcion")
	cabysCmd.MarkFlagsMutuallyExclusive("codigo", "top")
}
