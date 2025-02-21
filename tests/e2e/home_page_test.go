// tests/e2e/home_page_test.go

package e2e

import (
	"GoLang_FRT_E2E_Tests/pkg/pages"
	"testing"
	"log"
	"os"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	expectedTitle         = "Free Range Testers"
	expectedSectionsCount = 16
	expectedLinksCount    = 10
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

func verificarTitulo(page *pages.HomePage, t *testing.T) {
	startTime := time.Now()
	logger.Printf("🚀 Iniciando test de Título en FRT")
	logger.Printf("📡 Accediendo a la URL: %s", page.URL)
	titulo, err := page.GetTitle()
	if err != nil {
		t.Errorf("❌ Error obteniendo el título: %v", err)
		return
	}else{
		logger.Printf("📝 Título obtenido: %s", titulo)
		if assert.Equal(t, expectedTitle, titulo, "❌ Error Título no coincide") {
			logger.Printf("✅ Test de título completado en %.2f", time.Since(startTime).Seconds())
		}
		return
	}
}

func verificarSecciones(page *pages.HomePage, t *testing.T) {
	startTime := time.Now()
	logger.Printf("🚀 Iniciando test de secciones")
	secciones, err := page.GetSections()
	if err != nil {
		t.Errorf("❌ Error obteniendo secciones: %v", err)
		return
	}else{
		logger.Printf("📊 Número de secciones encontradas: %d", len(secciones))
		if assert.Len(t, secciones, expectedSectionsCount, "❌ Número de secciones no coincide") {
			logger.Printf("✅ Test de secciones completado en %.2f", time.Since(startTime).Seconds())
		}
		return
	}
}

func verificarEnlaces(page *pages.HomePage, t *testing.T) {
	startTime := time.Now()
	logger.Printf("🚀 Iniciando test de enlaces")
	enlaces, err := page.GetLinks()
	if err != nil {
		t.Errorf("❌ Error obteniendo enlaces: %v", err)
		return
	}else{
		if assert.Len(t, enlaces, expectedLinksCount, "❌ Número de enlaces no coincide") {
			logger.Printf("🔗 Número de enlaces encontrados: %d", len(enlaces))
			for i, enlace := range enlaces {
				logger.Printf("  🌐 Enlace %d: %s", i+1, enlace)
			}
			logger.Printf("✅ Test de enlaces completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func TestHomePage(t *testing.T) {
	page := pages.NewHomePage()
	t.Run("should have correct title", func(t *testing.T) {verificarTitulo(page, t)})
	t.Run("should have correct number of sections", func(t *testing.T) {verificarSecciones(page, t)})
	t.Run("should have correct number of links", func(t *testing.T) {verificarEnlaces(page, t)})
}