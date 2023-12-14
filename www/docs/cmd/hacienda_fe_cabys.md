## hacienda fe cabys

Permite obtener la información correspondiente al Catálogo de Bienes y Servicios (CABYS)

### Synopsis

Permite obtener la información correspondiente al Catalago de Bienes y Servicios (CABYS, https://www.hacienda.go.cr/docs/Catalogobienesyservicios.pdf),
a partir de la descripción de los bienes y servicios o su número de código.  

Puede utilizar los parámetros descripcion o codigo de la siguiente manera:

  - Por descripción del bien o servicio:
	
    hacienda fe cabys --descripcion="Jugo de tomate"
  
  - Por el código del bien o servicio:
	
    hacienda fe cabys --codigo="2132100000100"
  
  - El parámetro q puede ser usado en combinación con el parámetro top 
    para hacer una búsqueda limitada de bienes y servicios de la siguiente forma:
	
    hacienda fe cabys --descripcion="Jugo de tomate" --top=2


```
hacienda fe cabys [flags]
```

### Examples

```
  hacienda fe cabys --descripcion="Jugo de tomate"
  hacienda fe cabys --codigo="2132100000100"
  hacienda fe cabys --descripcion="Jugo de tomate" --top=2
```

### Options

```
  -c, --codigo string        Código del bien o servicio
  -d, --descripcion string   Descripción del bien o servicio
  -h, --help                 help for cabys
  -t, --top string           Limite de resultados para el parámetro descripción
  -v, --verbose              Verbose command
```

### SEE ALSO

* [hacienda fe](/hacienda/cmd/hacienda_fe/)	 - Muestra información relacionada con la Factura Electrónica.

