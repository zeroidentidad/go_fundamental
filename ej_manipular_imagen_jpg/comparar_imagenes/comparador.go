package main

import (
	"math"
	"sync"
)

//Compare func principal del archivo
func Compare(images []*image) chan Result {

	var wg sync.WaitGroup
	ch := make(chan Result)

	for _, needle := range images {
		for _, haystack := range images {

			if needle.width > haystack.width {
				continue
			}
			if needle.height > haystack.height {
				continue
			}
			if needle.name == haystack.name {
				continue
			}

			wg.Add(1)
			go func(n, h *image, ch chan Result) {
				defer wg.Done()
				comparePixels(n, h, ch)
			}(needle, haystack, ch)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

//comparar el píxel superior izquierdo de la aguja, n, con cada píxel en el pajar, h.
//recorriendo el slice de pixeles de la imagen
func comparePixels(n, h *image, ch chan Result) {

	var diff float64
	var wg sync.WaitGroup

	for i := 0; i < (h.width * h.height); i++ {

		// píxel actual en el pajar
		x := i % h.width
		y := i / h.width

		// la "aguja" debe encajar dentro del "pajar"
		if h.height-y < n.height {
			break
		}
		if h.width-x < n.width {
			continue
		}

		// encontrar la diferencia entre los píxeles
		diff = pixelDiff(n.pixels[0], h.pixels[i])
		if diff < limite {

			wg.Add(1)

			go func(n, h *image, i int, ch chan Result) {
				defer wg.Done()
				compareSequence(n, h, i, ch)
			}(n, h, i, ch)

			//fmt.Println("COMPARAR LANZADO:\t", n.name, h.name, "\t", diff, "\t", y, x) // DEBUGGING
		}
	}

	wg.Wait()
}

func pixelDiff(n, h pixel) float64 {
	var diff float64
	diff += math.Abs(float64(int(n.r) - int(h.r)))
	diff += math.Abs(float64(int(n.g) - int(h.g)))
	diff += math.Abs(float64(int(n.b) - int(h.b)))
	diff += math.Abs(float64(int(n.a) - int(h.a)))
	return diff
}

func compareSequence(n, h *image, hIdx int, ch chan Result) {

	var counter int
	var accumulator uint64
	hStartPix := hIdx

	for i := 0; i < n.height*n.width; i++ {

		diff := pixelDiff(n.pixels[i], h.pixels[hIdx])

		// cada fila debe tener un número mínimo de píxeles por debajo del umbral
		// (1) ¿la fila anterior tenía menos de 10?
		// (2) si está en una nueva fila, reinicie el contador
		// (3) si esta el píxel debajo del umbral, incrementar contador
		if ((i % n.width) == 0) && ((i / n.width) != 0) && counter < 10 {
			return
		}
		if (i % n.width) == 0 {
			counter = 0
		}
		if diff < limite {
			counter++
		}

		// si comienza una nueva fila, alinea pixeles
		if ((i + 1) % n.width) == 0 {
			hIdx += (h.width - n.width)
		}

		hIdx++
		accumulator += uint64(diff)
	}

	avgDiff := int(accumulator / uint64(n.height*n.width))
	ch <- Result{n, h, hStartPix, avgDiff}
}
