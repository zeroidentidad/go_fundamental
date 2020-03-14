
# Prueba contenedores locales modelos regresion

```shell
# ejecucion prueba contenedor modelo lineal 
docker run -v $PWD/training:/tmp/training -v $PWD/model:/tmp/model zeroidentidad/goregresiontrain:lineal ./goregresiontrain -inDir=/tmp/training -outDir=/tmp/model
#salida
Regression Formula:
Predicted = 152.1335 + bmi*949.4353

# ejecucion prueba contenedor modelo multiple
docker run -v $PWD/training:/tmp/training -v $PWD/model:/tmp/model zeroidentidad/goregresiontrain:multiple ./goregresiontrain -inDir=/tmp/training -outDir=/tmp/model
#salida
Regression Formula:
Predicted = 152.1335 + bmi*675.0698 + ltg*614.9505


# prueba ejecución prediccion
docker run -v $PWD/attributes:/tmp/attributes -v $PWD/model:/tmp/model zeroidentidad/goregresionpredict ./goregresionpredict -inModelDir=/tmp/model -inVarDir=/tmp/attributes -outDir=/tmp/model

# salida en model: 1.json, 2.json, 3.json, 4.json
```