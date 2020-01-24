
**Ahora que hay datos versionado en Pachyderm, es necesario saber cómo interactuar con esos datos. La forma principal es a través de pipelines de datos de Pachyderm. El mecanismo para interactuar con datos versionados cuando se usan canalizaciones, es un archivo de configuracion, para E/S.**

**Sin embargo, si queremos extraer manualmente ciertos conjuntos de datos versionados de Pachyderm, analizarlos interactivamente, entonces podemos usar la CLI de pachctl para obtener datos:**


### Obtener contenido del archivo manualmente:
```shell
pachctl get file repopruebas@master:blah.txt
```

### Descargar contenido del archivo manualmente:
```shell
pachctl get file repopruebas@master:blah.txt -o "Escritorio/blah.txt"
```


Referencia documentacion: https://docs.pachyderm.com/latest/reference/pachctl/pachctl_get_file/