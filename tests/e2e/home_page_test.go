// tests/e2e/home_page_test.go
package e2e

import (
	"GoLang_FRT_E2E_Tests/pkg/pages"
	"testing"
	"log"
	"os"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var logger *log.Logger

func TestMain(m *testing.M) {
	// Configurar el logger para escribir a un archivo
	logFile, err := os.OpenFile("../../reports/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo crear el archivo de log:", err)
	}
	defer logFile.Close()

	// Configurar el logger para incluir timestamp
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Ejecutar los tests
	code := m.Run()
	os.Exit(code)
}

const (
	expectedTitle         = "Free Range Testers"
	expectedSectionsCount = 16
	expectedLinksCount    = 10
)

func TestHomePage(t *testing.T) {
	t.Run("should have correct title", func(t *testing.T) {
		logger.Printf("🚀 Iniciando test de Título en FRT")
		startTime := time.Now()

		page := pages.NewHomePage()
		logger.Printf("📡 Accediendo a la URL: %s", page.URL)
		
		title, err := page.GetTitle()
		if err != nil {
			logger.Printf("❌ Error obteniendo el título: %v", err)
			t.Fatal(err)
		}

		logger.Printf("📝 Título obtenido: %s", title)
		assert.Equal(t, expectedTitle, title, "❌ Error obteniendo el título: No se encontró el título esperado")
		
		logger.Printf("✅ Test de título completado en %.2f", time.Since(startTime).Seconds())
		
		
	})

	
	t.Run("should have correct number of sections", func(t *testing.T) {
		logger.Printf("🚀 Iniciando test de secciones")
		startTime := time.Now()

		page := pages.NewHomePage()
		sections, err := page.GetSections()
		if err != nil {
			logger.Printf("❌ Error obteniendo las secciones: %v", err)
			t.Fatal(err)
		}

		assert.Equal(t, expectedSectionsCount, len(sections), "❌ Error obteniendo las secciones")
		logger.Printf("📊 Número de secciones encontradas: %d", len(sections))
		

		assert.Len(t, sections, expectedSectionsCount, "Number of sections doesn't match expected count")
		logger.Printf("✅ Test de secciones completado en %.2f", time.Since(startTime).Seconds())
	})

	
	t.Run("should have correct number of links", func(t *testing.T) {
		logger.Printf("🚀 Iniciando test de enlaces")
		startTime := time.Now()

		page := pages.NewHomePage()
		links, err := page.GetLinks()
		if err != nil {
			logger.Printf("❌ Error obteniendo los enlaces: %v", err)
			t.Fatal(err)
		}

		logger.Printf("🔗 Número de enlaces encontrados: %d", len(links))
		for i, link := range links {
			logger.Printf("  🌐 Enlace %d: %s", i+1, link)
		}

		assert.Len(t, links, expectedLinksCount, "Number of links doesn't match expected count")
		logger.Printf("✅ Test de enlaces completado en %.2f", time.Since(startTime).Seconds())
	})

	
	t.Run("should have valid structure", func(t *testing.T) {
		logger.Printf("🚀 Iniciando validación de estructura")
		startTime := time.Now()

		page := pages.NewHomePage()
		valid, err := page.VerifyStructure()
		if err != nil {
			logger.Printf("❌ Error verificando la estructura: %v", err)
			t.Fatal(err)
		}

		logger.Printf("🏗 Resultado de la validación: %v", valid)
		require.True(t, valid, "Homepage structure validation failed")
		logger.Printf("✅ Test de estructura completado en %.2f", time.Since(startTime).Seconds())
	})
}