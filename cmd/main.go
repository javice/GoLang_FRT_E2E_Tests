// cmd/main.go
package main

import (
	"log"
	"GoLang_FRT_E2E_Tests/pkg/pages"
)

func main() {
	// Crear instancia de la página
	homePage := pages.NewHomePage()

	// Verificar la estructura de la página
	valid, err := homePage.VerifyStructure()
	if err != nil {
		log.Fatalf("Error verificando la estructura: %v", err)
	}

	if valid {
		log.Println("La estructura de la página es válida")
	}

	// Obtener título
	title, err := homePage.GetTitle()
	if err != nil {
		log.Fatalf("Error obteniendo el título: %v", err)
	}
	log.Printf("Título de la página: %s", title)

	// Obtener secciones
	sections, err := homePage.GetSections()
	if err != nil {
		log.Fatalf("Error obteniendo las secciones: %v", err)
	}
	log.Printf("Número de secciones encontradas: %d", len(sections))

	// Obtener enlaces
	links, err := homePage.GetLinks()
	if err != nil {
		log.Fatalf("Error obteniendo los enlaces: %v", err)
	}
	log.Printf("Número de enlaces encontrados: %d", len(links))
}