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

type FEService service

func (s *FEService) ActividadEconomica(id string) (string, *http.Response, error) {

	u := fmt.Sprintf("fe/ae?identificacion=%v", id)

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

func (s *FEService) Exoneracion(authorizacion string) (string, *http.Response, error) {

	u := fmt.Sprintf("fe/ex?autorizacion=%v", authorizacion)

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
