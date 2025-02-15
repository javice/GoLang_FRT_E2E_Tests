//tests/e2e/sandbox_page_test.go

package e2e

import (
    "GoLang_FRT_E2E_Tests/pkg/pages"
    "testing"
    "time"
	

    "github.com/stretchr/testify/assert"
   "github.com/stretchr/testify/require"
)

const (
	expectedTitleSandbox  = "Automation Sandbox"
	expectedSectionsCountSandbox = 16
	expectedLinksCountSandbox    = 5
)
func TestSandboxPage(t *testing.T) {
    page := pages.NewSandboxPage()

	t.Run("should have correct title", func(t *testing.T) {
		logger.Printf("🚀 Iniciando test de Título en Sandbox de FRT")
		startTime := time.Now()

		//page := pages.NewHomePage()
		logger.Printf("📡 Accediendo a la URL: %s", page.URL)
		
		title, err := page.GetSandboxTitle()
		if err != nil {
			logger.Printf("❌ Error obteniendo el título: %v", err)
			t.Fatal(err)
		}

		logger.Printf("📝 Título obtenido: %s", title)
		assert.Equal(t, expectedTitleSandbox, title, "❌ Error obteniendo el título: No se encontró el título esperado")
		
		logger.Printf("✅ Test de título completado en %.2f", time.Since(startTime).Seconds())
		
		
	})
	
	t.Run("should have correct number of links", func(t *testing.T) {
		logger.Printf("🚀 Iniciando test de enlaces en Sandbox")
		startTime := time.Now()

		//page := pages.NewHomePage()
		links, err := page.GetSandboxLinks()
		if err != nil {
			logger.Printf("❌ Error obteniendo los enlaces: %v", err)
			t.Fatal(err)
		}

		logger.Printf("🔗 Número de enlaces encontrados: %d", len(links))
		for i, link := range links {
			logger.Printf("  🌐 Enlace %d: %s", i+1, link)
		}

		assert.Len(t, links, expectedLinksCountSandbox, "Number of links doesn't match expected count")
		logger.Printf("✅ Test de enlaces completado en %.2f", time.Since(startTime).Seconds())
	})

    t.Run("should click dynamic button", func(t *testing.T) {
        logger.Printf("🚀 Iniciando test de botón dinámico")
        startTime := time.Now()

        hiddenText,err := page.ClickDynamicButton()
        require.NoError(t, err, "❌ Error al hacer clic en el botón dinámico")

        logger.Printf("📝 Valor del Texto oculto: %s", hiddenText)
        logger.Printf("✅ Test de botón dinámico completado en %.2f", time.Since(startTime).Seconds())
    })

	
    t.Run("should insert text in textbox", func(t *testing.T) {
        logger.Printf("🚀 Iniciando test de cuadro de texto")
        startTime := time.Now()

        insertedText,err := page.InsertTextInTextbox("Texto de prueba")
        require.NoError(t, err, "❌ Error al insertar texto en el cuadro de texto")
		
        logger.Printf("📝 Valor del Textbox: %s", insertedText)
        logger.Printf("✅ Test de cuadro de texto completado en %.2f", time.Since(startTime).Seconds())
    })

    t.Run("should test checkboxes and radio buttons", func(t *testing.T) {
        logger.Printf("🚀 Iniciando test de checkboxes y radio buttons")
        startTime := time.Now()

        checkboxValue, radioValue, err := page.TestCheckboxesAndRadioButtons()
        require.NoError(t, err, "❌ Error al probar checkboxes y radio buttons")

        logger.Printf("📝 Valor del primer Checkbox: %s", checkboxValue)
        logger.Printf("📝 Valor del Radio Button: %s", radioValue)
        logger.Printf("✅ Test de checkboxes y radio buttons completado en %.2f", time.Since(startTime).Seconds())
    })

    t.Run("should click dropdowns", func(t *testing.T) {
        logger.Printf("🚀 Iniciando test de dropdowns")
        startTime := time.Now()

        firstDropdownValue, secondDropdownValue, err := page.ClickDropdowns()
        require.NoError(t, err, "Error al hacer clic en los dropdowns")

        logger.Printf("📝 Valor del primer Dropdown: %s", firstDropdownValue)
        logger.Printf("📝 Valor del segundo Dropdown : %s", secondDropdownValue)
        logger.Printf("✅ Test de dropdowns completado en %.2f", time.Since(startTime).Seconds())
    })

        
    
    t.Run("should handle popup", func(t *testing.T) {
        logger.Printf("🚀 Iniciando test de popup")
        startTime := time.Now()

        popupText, err := page.HandlePopup()
        require.NoError(t, err, "Error al manejar el popup")

        logger.Printf("📝 Texto del popup: %s", popupText)
        logger.Printf("✅ Test de popup completado en %.2f segundos", time.Since(startTime).Seconds())
    })
    
        
    
    t.Run("should interact with shadow DOM", func(t *testing.T) {
        logger.Printf("🚀 Iniciando test de Shadow DOM")
        startTime := time.Now()

        shadowContent, err := page.InteractWithShadowDOM()
        require.NoError(t, err, "Error al interactuar con el Shadow DOM")

        logger.Printf("✅ Contenido del Shadow DOM: %s", shadowContent)
        logger.Printf("✅ Test de Shadow DOM completado en %.2f", time.Since(startTime).Seconds())
    })
    t.Run("should interact with tables", func(t *testing.T) {
        logger.Printf("🚀 Iniciando test de tablas")
        startTime := time.Now()

        dynamicCellValueBefore, dynamicCellValueAfter, staticCellValueBefore, staticCellValueAfter, err := page.InteractWithTables()
        require.NoError(t, err, "Error al interactuar con las tablas")

        logger.Printf("📝 Valor de la celda dinámica antes: %s", dynamicCellValueBefore)
        logger.Printf("📝 Valor de la celda dinámica después: %s", dynamicCellValueAfter)
        logger.Printf("📝 Valor de la celda estática antes: %s", staticCellValueBefore)
        logger.Printf("📝 Valor de la celda estática luego: %s", staticCellValueAfter)

        if (dynamicCellValueBefore != dynamicCellValueAfter) && (staticCellValueBefore == staticCellValueAfter) {
            logger.Printf("✅ Test de tablas completado en %.2f segundos", time.Since(startTime).Seconds())
        } else {
            t.Errorf("❌ Los valores de las celdas no cumplen con las condiciones esperadas")
        }
    })
}