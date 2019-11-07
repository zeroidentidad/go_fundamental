# Cookies

Una cookie es un archivo que se almacena en la computadora del cliente.

Las cookies son escritas por dominios para almacenar información en la computadora del cliente.

## Dominio específico

El navegador solo envía las cookies al dominio que las escribió.

Con cada solicitud a un dominio específico, el navegador web del cliente busca para ver si hay una cookie de ese dominio en la computadora del cliente. Si hay una cookie que ha sido escrita por ese dominio en particular, entonces el navegador enviará la cookie con cada solicitud a ese dominio.

Las cookies son específicas del dominio.

## Limites

Podemos almacenar cualquier información que deseemos en una cookie hasta cierto límite de tamaño. El límite de tamaño de una cookie depende del navegador, pero generalmente es de alrededor de 4096 caracteres.

También hay un límite para la cantidad de cookies que puede escribir un dominio. Este límite también es específico del navegador.

Ver [este recurso](http://browsercookielimits.squawky.net/) para mas info.

## Caducar una cookie

Si los campos **Expires** o **MaxAge** no se establecen, entonces la cookie se elimina cuando se cierra el navegador. Esto se conoce coloquialmente como una "cookie de sesión".

Se puede caducar una cookie configurando uno de estos dos campos: **Expires** o **MaxAge**

**Expires** establece la hora exacta en que caduca la cookie. Expires esta **Deprecated**.

**MaxAge** establece cuánto tiempo debe vivir la cookie (en segundos).