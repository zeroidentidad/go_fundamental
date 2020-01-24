
### Crear archivo de prueba:
```shell
touch blah.txt
cat >> blah.txt
Este es un archivo de pruebas
```

### Para versionar el archivo en el repositorio, **repopruebas**, necesitamos hacer commit:
```shell
pachctl put file repopruebas@master -f blah.txt
```

### Confirmar que el archivo fue versionado en el repositorio:
```shell
pachctl list repo
```

### Ver el commit que se acaba de crear:
```shell
pachctl list commit repopruebas
```

### Ver el archivo en ese commit:
```shell
pachctl list file repopruebas@master
```

Referencia documentacion: https://docs.pachyderm.com/latest/getting_started/beginner_tutorial