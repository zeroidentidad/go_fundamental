
# Verificar y aplicar los tipos esperados:

Esto puede parecer obvio, pero con demasiada frecuencia se pasa por alto cuando se usan lenguajes de tipado dinamico. Aunque es un poco detallado, analizar explícitamente los datos en los tipos esperados y manejar los errores relacionados puede ahorrar grandes dolores de cabeza en el futuro.

# Estandarizar y simplificar entrada/salida de datos:

Existen muchos paquetes de terceros para manejar ciertos tipos de datos o interacciones con ciertas fuentes de datos. Sin embargo, si se estandariza las formas en que interactúa con las fuentes de datos, particularmente centradas en el uso de stdlib (standar library), puede desarrollar patrones predecibles y mantener la coherencia dentro del equipo de trabajo. Un buen ejemplo de esto es la opción de utilizar database/sql para las interacciones de la base de datos en lugar de utilizar varias API y DSL de terceros.

# Versionar los datos:

Los modelos de aprendizaje automático producen resultados extremadamente diferentes dependiendo de los datos de entrenamiento que use, su elección de parámetros y datos de entrada. Por lo tanto, es imposible reproducir resultados sin versionar tanto el código como los datos.