# Formatear con tiempo

Como MST es GMT-0700, el tiempo de referencia puede considerarse como:

01/02 03:04:05PM '06 -0700

(Enero 02, 2006)

[godoc.org/time](https://godoc.org/time#pkg-constants)

***

## func (Time) Format

```go
func (t Time) Format(layout string) string
```

Format devuelve una representación textual del valor de tiempo formateado de acuerdo con el diseño, que define el formato mostrando cómo se define el tiempo de referencia

```go
Mon Jan 2 15:04:05 -0700 MST 2006
```

se mostraría si fuera el valor; Sirve como un ejemplo de la salida deseada. Las mismas reglas de visualización se aplicarán al valor de tiempo.

Un segundo fraccionario se representa agregando un punto y ceros al final de la sección de segundos de la cadena de diseño, como en "15: 04: 05.000" para formatear una marca de tiempo con una precisión de milisegundos.

Los diseños predefinidos ANSIC, UnixDate, RFC3339 y otros describen representaciones estándar y convenientes del tiempo de referencia. Para obtener más información sobre los formatos y la definición del tiempo de referencia, consulte la documentación de ANSIC y las otras constantes definidas por este paquete.

# MST

La zona horaria de montaña de América del Norte mantiene el tiempo restando siete horas del tiempo universal coordinado (UTC), durante los días más cortos de otoño e invierno (UTC − 7), y restando seis horas durante el horario de verano en primavera, verano, y principios de otoño (UTC − 6). La hora del reloj en esta zona se basa en la hora solar media en el meridiano 105 al oeste del Observatorio de Greenwich. En los Estados Unidos, la especificación exacta para la ubicación de zonas horarias y las líneas divisorias entre zonas se establece en el Código de Regulaciones Federales en 49 CFR 71.

# UTC

El tiempo universal coordinado, abreviado como UTC, es el estándar de tiempo primario por el cual el mundo regula los relojes y el tiempo. Está dentro de aproximadamente 1 segundo del tiempo solar medio a 0 ° de longitud; no observa el horario de verano. Es uno de los varios sucesores estrechamente relacionados con el tiempo medio de Greenwich (GMT). Para la mayoría de los propósitos, UTC se considera intercambiable con GMT, pero GMT ya no se define con precisión por la comunidad científica.

# GMT

El tiempo medio de Greenwich (GMT) es el tiempo solar medio en el Observatorio Real de Greenwich, Londres.

GMT se usaba anteriormente como el estándar internacional de tiempo civil, ahora reemplazado en esa función por el Tiempo Universal Coordinado (UTC).