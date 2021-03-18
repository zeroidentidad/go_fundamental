package web

import (
	"pastein/pkg/models"
	"path/filepath"
	"text/template"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// filepath.Glob para obtener slice de todas las rutas de archivo con
	// la extensión '.page.tmpl'. Esto esencialmente da una slice de todos los
	// templates de 'page' para la aplicación.
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extraer nombre del archivo (como 'home.page.tmpl') de ruta completa del archivo
		// y asignar a la variable de name.
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// Agrega el conjunto de templates a la caché, usando el nombre de la página
		// (como 'home.page.tmpl') como clave.
		cache[name] = ts
	}

	return cache, nil
}
