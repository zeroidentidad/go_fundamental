# Session

Como se crea estado:

- Almacenar una identificación única en la cookie.

- En nuestro servidor, asociar a cada usuario con una identificación única.

- Esto permitirá identificar a cada usuario que visite nuestro webasite.

## Seguridad

Hay dos factores que contribuyen a la seguridad de una sesión creada usando una cookie y una ID única:

1. Unicidad del ID
1. Cifrado en transmision con HTTPS

Se puede usar cualquier ID única que se desee: un identificador único universal [(UUID)](https://en.wikipedia.org/wiki/Universally_unique_identifier) o una clave de base de datos. Si se está utilizando una clave de base de datos, asegúrarse de que no sea la clave para el usuario sino una tabla de sesión separada.

Un UUID es totalmente unico. " ... solo después de generar mil millones de UUID por segundo durante los próximos 100 años, la probabilidad de crear solo un duplicado sería de aproximadamente el 50%."

Un UUID no puede ser interceptado en transmision si estamos usando HTTPS.