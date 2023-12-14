## hacienda fe ae

Permite consultar información relacionada al contribuyente

### Synopsis

Permite obtener el nombre, el tipo de identificación, el régimen, 
la situación tributaria y las actividades económicas asociadas a un contribuyente, 
usando como parámetro el número de identificación (sin hacer uso de guiones).

Identificaciones soportadas:

  1. Físicas nacionales
  2. Jurídicas nacionales
  3. DIMEX: Documento de Identificación Migratorio para Extranjeros
  4. NITE: Numero de Identificación Tributario Especial

Restricciones:

  - Físicas: debe contener 9 números, sin cero al inicio y sin guiones.
  - DIMEX: debe contener 11 o 12 dígitos, sin cero al inicio y sin guiones.
  - Jurídicas y NITE: debe contener 10 dígitos, sin guiones y sin utilizar los dos últimos dígitos verificadores.

Ver también:

  https://atv.hacienda.go.cr/ATV/frmConsultaSituTributaria.aspx


```
hacienda fe ae [flags]
```

### Options

```
  -h, --help        help for ae
  -i, --id string   Número de identificación del contribuyente (default "0")
  -v, --verbose     Verbose command
```

### SEE ALSO

* [hacienda fe](/hacienda/cmd/hacienda_fe/)	 - Muestra información relacionada con la Factura Electrónica.

