package stringutil

func reverseDos(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// esto demuestra cómo una función no exportada
// puede ser usado por una función exportada en el mismo paquete
