//tests/e2e/sandbox_page_test.go

package e2e

import (
    "GoLang_FRT_E2E_Tests/pkg/pages"
    "testing"
    "time"
    
    "github.com/stretchr/testify/assert"
   //"github.com/stretchr/testify/require"
)

const (
	expectedTitleSandbox  = "Automation Sandbox"
	expectedSectionsCountSandbox = 16
	expectedLinksCountSandbox    = 5
)
func verificarTituloSandbox(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de Título en Sandbox de FRT")
    logger.Printf("📡 Accediendo a la URL: %s", page.URL)
    titulo, err := page.GetSandboxTitle()
    if err != nil {
        t.Errorf("❌ Error obteniendo el título: %v", err)
        return
    }else{
        logger.Printf("📝 Título obtenido: %s", titulo)
        if assert.Equal(t, expectedTitleSandbox, titulo, "❌ Error Título no coincide") {
            logger.Printf("✅ Test de título completado en %.2f", time.Since(startTime).Seconds())
        }
        return
    }
}

func verificarEnlacesSandbox(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de enlaces en Sandbox")
    enlaces, err := page.GetSandboxLinks()
    if err != nil {
        t.Errorf("❌ Error obteniendo enlaces: %v", err)
        return
    }else{
        if assert.Len(t, enlaces, expectedLinksCountSandbox, "❌ Número de enlaces no coincide") {
            logger.Printf("🔗 Número de enlaces encontrados: %d", len(enlaces))
            for i, enlace := range enlaces {
                logger.Printf("  🌐 Enlace %d: %s", i+1, enlace)
            }
            logger.Printf("✅ Test de enlaces completado en %.2f", time.Since(startTime).Seconds())
        }
        return
    }
}

func verificarBotonDinamico(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de botón dinámico en Sandbox")
    boton, err := page.ClickDynamicButton()
    if err != nil {
        t.Errorf("❌ Error obteniendo el botón dinámico: %v", err)
        return
    }else{
        logger.Printf("📝 Valor del botón dinámico: %s", boton)
        logger.Printf("✅ Test de botón dinámico completado en %.2f", time.Since(startTime).Seconds())
        return
    }
}

func verificarTextbox(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de textbox en Sandbox")
    textbox, err := page.InsertTextInTextbox("Texto de prueba")
    if err != nil {
        t.Errorf("❌ Error obteniendo el textbox: %v", err)
        return
    }else{
        logger.Printf("📝 Valor del textbox: %s", textbox)
        logger.Printf("✅ Test de textbox completado en %.2f", time.Since(startTime).Seconds())
        return
    }
}

func verificarCheckboxesRadioButtons(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de checkboxes en Sandbox")
    checkboxes, radioValue,err := page.TestCheckboxesAndRadioButtons()
    if err != nil {
        t.Errorf("❌ Error obteniendo los checkboxes: %v", err)
        return
    }else{
        logger.Printf("📝 Valores de los checkboxes: %s", checkboxes)
        logger.Printf("📝 Valor del Radio Button: %s", radioValue)
        logger.Printf("✅ Test de checkboxes y radio buttons completado en %.2f", time.Since(startTime).Seconds())
        return
    }
}

func verificarDropdowns(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de dropdowns en Sandbox")
    primerDropdown, segundoDropdown, err := page.ClickDropdowns()
    if err != nil {
        t.Errorf("❌ Error obteniendo los dropdowns: %v", err)
        return
    }else{
        logger.Printf("📝 Valor del primer Dropdown: %s", primerDropdown)
        logger.Printf("📝 Valor del segundo Dropdown : %s", segundoDropdown)
        logger.Printf("✅ Test de dropdowns completado en %.2f", time.Since(startTime).Seconds())
        return
    }
}

func verificarPopup(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de popup en Sandbox")
    popup, err := page.HandlePopup()
    if err != nil {
        t.Errorf("❌ Error obteniendo el popup: %v", err)
        return
    }else{
        logger.Printf("📝 Valor del popup: %s", popup)
        logger.Printf("✅ Test de popup completado en %.2f", time.Since(startTime).Seconds())
        return
    }
}

func verificarShadowDom(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de Shadow DOM en Sandbox")
    shadowDom, err := page.InteractWithShadowDOM()
    if err != nil {
        t.Errorf("❌ Error obteniendo el Shadow DOM: %v", err)
        return
    }else{
        logger.Printf("📝 Valor del Shadow DOM: %s", shadowDom)
        logger.Printf("✅ Test de Shadow DOM completado en %.2f", time.Since(startTime).Seconds())
        return
    }
}

func verificarTablas(page *pages.SandboxPage, t *testing.T) {
    startTime := time.Now()
    logger.Printf("🚀 Iniciando test de tablas en Sandbox")
    valorCeldaDinamicaAntes, valorCeldaDinamicaDespues, valorCeldaEstaticaAntes, valorCeldaEstaticaDespues, err := page.InteractWithTables()
    if err != nil {
        t.Errorf("❌ Error obteniendo las tablas: %v", err)
        return
    }else{
        logger.Printf("📝 Valor de la celda dinámica antes: %s", valorCeldaDinamicaAntes)
        logger.Printf("📝 Valor de la celda dinámica después: %s", valorCeldaDinamicaDespues)
        logger.Printf("📝 Valor de la celda estática antes: %s", valorCeldaEstaticaAntes)
        logger.Printf("📝 Valor de la celda estática luego: %s", valorCeldaEstaticaDespues)
        logger.Printf("✅ Test de tablas completado en %.2f", time.Since(startTime).Seconds())
        if (valorCeldaDinamicaAntes != valorCeldaDinamicaDespues) && (valorCeldaEstaticaAntes == valorCeldaEstaticaDespues) {
            logger.Printf("✅ Test de tablas completado en %.2f segundos", time.Since(startTime).Seconds())
        } else {
            t.Errorf("❌ Los valores de las celdas no cumplen con las condiciones esperadas")
        }
        return
    }
}

func TestSandboxPage(t *testing.T) {
    page := pages.NewSandboxPage()
    t.Run("should have correct title", func(t *testing.T){verificarTituloSandbox(page, t)})
    //t.Run("should have correct number of sections", func(t *testing.T){verificarSeccionesSandbox(page, t)})
    t.Run("should have correct number of links", func(t *testing.T){verificarEnlacesSandbox(page, t)})
    t.Run("should click dynamic button", func(t *testing.T){verificarBotonDinamico(page, t)})
    t.Run("should insert text in textbox", func(t *testing.T){verificarTextbox(page, t)})
    t.Run("should test checkboxes and radio buttons", func(t *testing.T){verificarCheckboxesRadioButtons(page, t)})
    t.Run("should click dropdowns", func(t *testing.T){verificarDropdowns(page, t)})
    t.Run("should handle popup", func(t *testing.T){verificarPopup(page, t)})
    t.Run("should interact with shadow DOM", func(t *testing.T){verificarShadowDom(page, t)})
    t.Run("should interact with tables", func(t *testing.T){verificarTablas(page, t)})
}
