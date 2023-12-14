## hacienda fe agropecuario

Permite consultar información relacionada al productor agropecuario

### Synopsis

Permite obtener el nombre, el estado, el indicador de si está activo
y la fecha de baja de los productores agropecuarios en el MAG, usando como 
parámetro el número de "identificacion" (sin hacer uso de guiones). 

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
hacienda fe agropecuario [flags]
```

### Options

```
  -h, --help        help for agropecuario
  -i, --id string   Número de identificación del productor agropecuario (default "0")
  -v, --verbose     Verbose command
```

### SEE ALSO

* [hacienda fe](/cmd/hacienda_fe/)	 - Muestra información relacionada con la Factura Electrónica.

