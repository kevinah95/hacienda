/*
 * Copyright (c) 2023 Kevin Hernández Rostrán
 * Use of this source code is governed by the Apache 2.0 license that can be found in the LICENSE file.
 */

package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type INService service

func (s *INService) TipoDeCambio(typeOfExchange string, from string, to string) (string, *http.Response, error) {
	u := ""
	switch typeOfExchange {
	case "dolar":
		if from != "" && to != "" {
			u = fmt.Sprintf("indicadores/tc/dolar/historico?d=%v&h=%v", from, to)
		} else {
			u = "indicadores/tc/dolar"
		}
	case "euro":
		u = "indicadores/tc/euro"
	default:
		u = "indicadores/tc"
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return "", nil, err
	}

	response, err := s.client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData), response, nil

}
