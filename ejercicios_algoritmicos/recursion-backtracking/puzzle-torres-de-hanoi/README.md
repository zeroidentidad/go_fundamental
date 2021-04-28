
### Solución:

Las Torres de Hanoi son un puzzle matemático. Consiste en tres varillas (o clavijas o torres) y varios discos de diferentes tamaños que pueden deslizarse sobre cualquier varilla. El puzzle comienza con los discos en una varilla en orden ascendente de tamaño, el más pequeño en la parte superior, creando así una forma cónica. El objetivo del puzzle es mover toda la pila a otra barra, cumpliendo las siguientes reglas:

- Solo se puede mover un disco a la vez.

- Cada movimiento consiste en tomar el disco superior de una de las varillas y deslizarlo sobre otra varilla, encima de los otros discos que ya pueden estar presentes en esa varilla.

- No se puede colocar ningún disco encima de un disco más pequeño.

### Algoritmo:

- Mover los **n-1** discos superiores de la fuente a la torre auxiliar,

- Mover el **n°** disco de la fuente a la torre destino,

- Mover **n-1** discos de la torre auxiliar a la torre destino,

- La transferencia de los **n-1** discos superiores desde la fuente a la torre auxiliar puede considerarse nuevamente como un problema nuevo y puede resolverse de la misma manera. Una vez que se resuelve Torres de Hanoi con tres discos, puede resolverse con cualquier cantidad de discos con el algoritmo anterior.