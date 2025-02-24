// tests/e2e/avis_page_test.go
package e2e

import (
	"testing"
	"time"
	//"fmt"

	"GoLang_FRT_E2E_Tests/pkg/pages"
	"github.com/stretchr/testify/require"
	
)

const (
	urlAvis = "https://www.avis.es"
	expectedTitleAvis  = "Alquiler de coches en España, Europa y resto del mundo | Avis"
	expectedSectionsCountAvis = 16
	expectedLinksCountAvis    = 5
	pickupLocation = "Madrid-Barajas Adolfo Suárez T1 y T4 - ESP"
	returnLocation = "Barcelona-El Prat T1 y T2 - ESP"
)

func verificarBusquedaAvis(page *pages.AvisPage, t *testing.T) {
	startTime := time.Now()
    logger.Printf("🚀 Iniciando test de AVIS")
    logger.Printf("📡 Accediendo a la URL: %s",urlAvis )
	require.NoError(t, page.NavigateTo("https://www.avis.es"))
	require.NoError(t, page.AcceptCookies())
	require.NoError(t, page.SearchVehicles(time.Now(), time.Now(), pickupLocation, returnLocation))
	
	// Verificar que el título de la página de coches disponibles contiene el texto esperado
    expectedTitle := "Resultados Búsqueda"
	actualTitle, err := page.Title()
    require.NoError(t, err)
    require.Contains(t, actualTitle, expectedTitle, "El título de la página de resultados no es el esperado")
	logger.Printf("📝 Título obtenido: %s", actualTitle)

	// Verificar que se han encontrado vehículos
	vehicles, err := page.AvailableVehicles()
	require.NoError(t, err)
	require.Greater(t, len(vehicles), 0, "❌ No se han encontrado vehículos disponibles")
	logger.Printf("🚗 Vehículos encontrados: %d", len(vehicles))
	
	logger.Printf("✅ Test de búsqueda completado en %.2f", time.Since(startTime).Seconds())
}


func TestAvisPage(t *testing.T) {
    avisPage := pages.NewAvisPage()
    defer avisPage.Close()

    t.Run("should search for a vehicle", func(t *testing.T) {
        verificarBusquedaAvis(avisPage, t)
    })
}
