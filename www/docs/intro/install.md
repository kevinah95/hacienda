# Instalación

Puede instalar el binario precompilado (de varias maneras) o compilar
desde el código fuente.

A continuación puedes encontrar los pasos para cada una de ellas.

## Instalar el binario precompilado

### [goblin.run](https://goblin.run)

```sh
$ curl -sf http://goblin.run/github.com/kevinah95/hacienda | sh
```

### Homebrew tap

```
$ brew install kevinah95/tap/hacienda
```

### Go install

```bash
go install github.com/kevinah95/hacienda@latest
```

Requiere Go 1.20.

### Manualmente

- Descargar el binario de [releases](https://github.com/kevinah95/hacienda/releases) 
y copiarlos en la ubicación deseada.

## Verificar los binarios

Este proyecto utiliza Cosign para firmar los archivos binarios. 

!!! note "¿Qué es Cosign?"
    Cosign es una utilidad de línea de comandos que puede firmar y verificar artefactos de software, como imágenes de contenedores y blobs.

Para verificar la firma con `cosign`, vea este ejemplo:

```bash
$ cosign verify-blob --key cosign.pub --signature signatures/hacienda.sig  ./hacienda
```

- `cosign.pub`: Se encuentra en el directorio raíz del proyecto.
- `hacienda.sig`: Se encuentra en el archivo comprimido en [releases](https://github.com/kevinah95/hacienda/releases).
- `./hacienda`: Es el artefacto de software a verificar.

Para más información visite la [documentación de sigstore](https://docs.sigstore.dev) o la sección de [firma de GoReleaser](https://goreleaser.com/customization/sign).

## Compilar desde el código fuente

Aquí tienes dos opciones:

Si quieres contribuir al proyecto, sigue los
pasos en nuestra [guía de contribución](/hacienda/contributing).

Si, por cualquier motivo, quieres compilar desde el código fuente, sigue estos pasos:

**clonar:**

```bash
git clone https://github.com/goreleaser/goreleaser
cd goreleaser
```

**obtener las dependencias:**

```bash
go mod tidy
```

**compilar:**

```bash
go build -o goreleaser .
```

**comprobar que funciona:**

```bash
./goreleaser --version
```