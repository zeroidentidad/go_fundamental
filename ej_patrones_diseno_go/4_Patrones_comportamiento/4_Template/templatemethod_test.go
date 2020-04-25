package templatemethod

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	deployAndroid := DeployAndroid{Deploy{}}
	deployAndroid.Construir(&deployAndroid)

	fmt.Print("\n")

	deployiOS := DeployiOS{Deploy{}}
	deployiOS.Construir(&deployiOS)
}
