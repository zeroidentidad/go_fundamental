package proxy

// Sujeto Interface
type NavegadorInterface interface {
	Direccion(string) string
}

// Sujeto Real
type Navegador struct{}

func (n *Navegador) Direccion(url string) string {
	return "Respuesta de la url " + url
}

// Proxy
type NavegadorProxy struct {
	navegador NavegadorInterface
}

func (n *NavegadorProxy) Direccion(url string) string {
	if url == "http://twitter.com" || url == "http://facebook.com" {
		return "Acceso restringido a " + url
	}

	return n.navegador.Direccion(url)
}
