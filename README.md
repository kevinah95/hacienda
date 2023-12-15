# Hacienda CLI

[![Made in Costa Rica](https://img.shields.io/badge/made%20in-%20Costa%20Rica-blue.svg?logo=data:image/svg%2bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGlkPSJmbGFnLWljb25zLWNyIiB2aWV3Qm94PSIwIDAgNjQwIDQ4MCI+CiAgPGcgZmlsbC1ydWxlPSJldmVub2RkIiBzdHJva2Utd2lkdGg9IjFwdCI+CiAgICA8cGF0aCBmaWxsPSIjMDAwMGI0IiBkPSJNMCAwaDY0MHY0ODBIMHoiLz4KICAgIDxwYXRoIGZpbGw9IiNmZmYiIGQ9Ik0wIDc1LjRoNjQwdjMyMi4zSDB6Ii8+CiAgICA8cGF0aCBmaWxsPSIjZDkwMDAwIiBkPSJNMCAxNTcuN2g2NDB2MTU3LjdIMHoiLz4KICA8L2c+Cjwvc3ZnPgo=)](https://es.wikipedia.org/wiki/Costa_Rica)
[![GitHub license](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://www.apache.org/licenses/LICENSE-2.0)
[![GitHub release (with filter)](https://img.shields.io/github/v/release/kevinah95/hacienda)](https://github.com/kevinah95/hacienda/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevinah95/hacienda)](https://goreportcard.com/report/github.com/kevinah95/hacienda)
[![Go Reference](https://pkg.go.dev/badge/github.com/kevinah95/hacienda.svg)](https://pkg.go.dev/github.com/kevinah95/hacienda)


Hacienda es una aplicación que se conecta al API del Ministerio de Hacienda de Costa Rica.

<img src="./examples/hacienda.gif" alt="Ejemplo hacienda" />

## Instalación

### Vía [goblin.run](https://goblin.run)

```sh
$ curl -sf http://goblin.run/github.com/kevinah95/hacienda | sh
```

### Homebrew tap

```
$ brew install kevinah95/tap/hacienda
```

### Go install

```
$ go install github.com/kevinah95/hacienda@latest
```

### Manual

- Descargar el binario de [releases](https://github.com/kevinah95/hacienda/releases).

## Verificar

Este proyecto utiliza Cosign para firmar los archivos binarios. 

> Cosign es una utilidad de línea de comandos que puede firmar y verificar artefactos de software, como imágenes de contenedores y blobs.

Para verificar la firma con `cosign`, vea este ejemplo:

```bash
$ cosign verify-blob --key cosign.pub --signature signatures/hacienda.sig  ./hacienda
```

- `cosign.pub`: Se encuentra en el directorio raíz del proyecto.
- `hacienda.sig`: Se encuentra en el archivo comprimido en [releases](https://github.com/kevinah95/hacienda/releases).
- `./hacienda`: Es el artefacto de software a verificar.

Para más información visite la [documentación de sigstore](https://docs.sigstore.dev) o la sección de [firma de GoReleaser](https://goreleaser.com/customization/sign).

## Licencia

```
Copyright 2023 Kevin Hernández Rostrán

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```