
### Problema 1:

1- Generar todas las cadenas de *n* bits. Suponer que **A[0...n − 1]** es un array de tamaño *n*.

### Solución:

Sea *T(n)* el tiempo de ejecución de *binary(n)*. Suponer que la función *printf* tarda un tiempo *O(1)*.

T(n)= {c, if n<0 ... 2T(n-1)+d, de lo contrario}

Código en **solucion-1.go**

Usando el teorema de Resta y Conquista del Maestro se obtiene: *T(n) = O(2n)*. Esto significa que el algoritmo para generar cadenas de bits es óptimo.