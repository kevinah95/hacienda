## hacienda fe ex

Permite consultar la información correspondiente a una exoneración

### Synopsis

Permite obtener la información correspondiente a una exoneración.
Requiere el parámetro "autorizacion" cuyo formato debe seguir la regla "al-00000000-00".

Cuando una autorización posee CABYS asociados, el campo "poseeCabys" tendrá un valor "true" y en consecuencia 
aparecerá el campo "cabys" que corresponde a un array de códigos CABYS. 
Si el valor de "poseeCabys" es "false", la respuesta no incluye el arreglo de códigos.

Restricciones:

  - Si el documento de exoneración no existe se devuelve el código "404".  
  - Si el formato del parámetro "autorizacion" es incorrecto o no existe como parámetro, 
    se devuelve el código "400".




```
hacienda fe ex [flags]
```

### Options

```
  -a, --authz string   Número de documento. (default "al-00000000-00")
  -h, --help           help for ex
  -v, --verbose        verbose command
```

### SEE ALSO

* [hacienda fe](/cmd/hacienda_fe/)	 - Muestra información relacionada con la Factura Electrónica.

