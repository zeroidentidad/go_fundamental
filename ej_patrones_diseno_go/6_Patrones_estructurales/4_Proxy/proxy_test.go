package proxy

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	navegadorProxy := &NavegadorProxy{&Navegador{}}

	fmt.Printf("%s\n", navegadorProxy.Direccion("http://google.com"))
	fmt.Printf("%s\n", navegadorProxy.Direccion("http://twitter.com"))
	fmt.Printf("%s\n", navegadorProxy.Direccion("http://facebook.com"))
}
