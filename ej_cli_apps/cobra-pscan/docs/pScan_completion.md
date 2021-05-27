## pScan completion

Generar completado de bash para comandos

### Sinopsis

Para cargar completado ejecutar

source <(pScan completion)

Para cargar completado automáticamente al abrir terminal, agregar esta línea a su archivo .bashrc

$ ~/.bashrc

source <(pScan completion)

```
pScan completion [flags]
```

### Opciones

```
  -h, --help   ayuda para completado
```

### Opciones heredadas de los comandos principales

```
      --config string       archivo configuración (default $HOME/.pScan.yaml)
  -f, --hosts-file string   pScan archivo hosts (default "pScan.hosts")
```

### VER TAMBIÉN

* [pScan](pScan.md)	 - Escáner de puerto TCP