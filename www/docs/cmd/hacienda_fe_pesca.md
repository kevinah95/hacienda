## hacienda fe pesca

Permite consultar información relacionada al productor pesquero

### Synopsis

Permite obtener los siguientes datos:

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


```
hacienda fe pesca [flags]
```

### Options

```
  -h, --help        help for pesca
  -i, --id string   Número de identificación del productor pesquero (default "0")
  -v, --verbose     Verbose command
```

### SEE ALSO

* [hacienda fe](/cmd/hacienda_fe/)	 - Muestra información relacionada con la Factura Electrónica.

