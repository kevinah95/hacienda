# Contribuír

Al participar en este proyecto, acepta cumplir nuestro
[código de conducta](https://github.com/kevinah95/hacienda/blob/main/CODE_OF_CONDUCT.md).

## Configuración

`hacienda` está escrito en [Go](https://golang.org/).

Prerrequisitos:

- [Task](https://taskfile.dev/installation)
- [Go 1.20+](https://go.dev/doc/install)

Clonar `hacienda`:

```sh
git clone git@github.com:kevinah95/hacienda.git
```

`cd` en el directorio e instalar las dependencias:

```sh
task setup
```

## Pruebe su cambio

Puedes crear una rama para sus cambios e intentar compilar desde el código fuente sobre la marcha:

```sh
task build
```

Cuando esté satisfecho con los cambios, le sugerimos que ejecute:

```sh
task ci
```

## Commits

Los mensajes Commit deben estar bien formateados, y para que eso sea "estandarizado", 
estamos usando la especificación [Commits Convencionales v1.0.0](https://www.conventionalcommits.org/es/v1.0.0).

Seguir esta convención nos permite ofrecer un proceso de publicación automatizado que también genera un registro de cambios (CHANGELOG) detallado.

## Pull Requests

Haga push de su rama a su fork de `hacienda` y abra un pull request contra la rama principal.

## Releases

Los releases son gestionados por [Release Please](https://github.com/googleapis/release-please) que se ejecuta en una acción de GitHub cada vez que se hace un commit en la rama `main`.

Release Please analiza los [commits](#commits) y abre (o actualiza) un pull request contra la rama `main` que contiene actualizaciones de los releases y el Changelog dentro del proyecto. Si no detecta ningún commit de breaking change, sólo incrementará la versión "patch"; sin embargo, si detecta un commit de breaking change, incrementará el número de versión "minor" para indicar un breaking release.

Cuando estemos listos para liberar la versión, aprobamos y hacemos squash del pull request del release en `main`. **Release Please** detectará este merge y generará las etiquetas apropiadas para la versión. Puede ser que se activen pasos adicionales dentro de la Acción GitHub para automatizar otras partes del proceso de publicación.

## Changelog

El Changelog de `hacienda` es gestionado automáticamente por Release Please a partir de los [Commits Convencionales](#commits) (como se discutió anteriormente).