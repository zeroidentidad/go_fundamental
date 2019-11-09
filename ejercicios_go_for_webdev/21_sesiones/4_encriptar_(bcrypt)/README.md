Dicen los sabios informaticos que nunca se debe almacenar una contraseña sin cifrarla.

Usaremos ```bcrypt``` para encriptar contraseñas.

# Paso 1: 

Se tendrá que obtener este paquete:

```
go get golang.org/x/crypto/bcrypt
```

# Paso 2:
Cambiar el campo de contraseña de la estructura de usuario a los datos type []byte

# Paso 3:
Usar bcrypt para cifrar contraseña antes de guardarla.