# Plantillas anidadas

[documentación de plantillas anidadas](https://godoc.org/text/template#hdr-Nested_template_definitions)

## define: 
``` Go
{{define "TemplateName"}}
inserte contenido aquí
{{end}}
```
## use: 
``` Go
{{template "TemplateName"}}
```