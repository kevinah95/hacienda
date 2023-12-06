/*
 * Copyright (c) 2023 Kevin Hernández Rostrán
 * Use of this source code is governed by the Apache 2.0 license that can be found in the LICENSE file.
 */

package cmd

import (
	"github.com/spf13/cobra"
)

// feCmd represents the fe command
var feCmd = &cobra.Command{
	Use:     "fe",
	Aliases: []string{"facturaelectronica"},
	Short:   "Muestra información relacionada con la Factura Electrónica.",
	Long: `La Factura Electrónica es un comprobante electrónico para respaldar la
venta de bienes y la prestación de servicios, que debe ser generado y 
transmitido en el mismo acto de la compraventa o la prestación del 
servicio.

Más información sobre la Factura Electrónica en:

	https://www.hacienda.go.cr/docs/AspectosGeneralesComprobantesElectronicos.pdf
`,
}

func init() {
	rootCmd.AddCommand(feCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// feCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// feCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
