// pkg/pages/sandbox_page.go
package pages

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)
// PageError representa los errores personalizados para el scraping
type SandboxPageError struct {
	Message string
	Err     error
}

func (e *PageError) SandboxError() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// SandboxPage representa la página del sandbox de FRT
type SandboxPage struct {
	URL    string
	client *http.Client
}

// NewSandboxPage crea una nueva instancia de HomePage
func NewSandboxPage() *SandboxPage {
	return &SandboxPage{
		URL: "https://thefreerangetester.github.io/sandbox-automation-testing/",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// fetchSandboxContent obtiene el contenido de la página
func (h *SandboxPage) fetchSandboxContent() (*goquery.Document, error) {
	req, err := http.NewRequest("GET", h.URL, nil)
	if err != nil {
		return nil, &PageError{"Error creating request", err}
	}

	req.Header.Set("User-Agent", "FreeRangeTesters E2E Tests")
	
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, &PageError{"Error fetching page", err}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &PageError{
			Message: fmt.Sprintf("Status code error: %d", resp.StatusCode),
			Err:     nil,
		}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, &PageError{"Error parsing HTML", err}
	}

	return doc, nil
}

// GetTitle obtiene el título de la página
func (h *SandboxPage) GetSandboxTitle() (string, error) {
	doc, err := h.fetchSandboxContent()
	if err != nil {
		return "", err
	}

	title := doc.Find("title").Text()
	if title == "" {
		return "", &PageError{"Title not found", nil}
	}

	return strings.TrimSpace(title), nil
}

// GetSections obtiene todas las secciones de la página
func (h *SandboxPage) GetSandboxSections() ([]string, error) {
	doc, err := h.fetchSandboxContent()
	if err != nil {
		return nil, err
	}

	var sections []string
	doc.Find("[id^='page_section']").Each(func(_ int, s *goquery.Selection) {
		sections = append(sections, s.Text())
	})

	if len(sections) == 0 {
		return nil, &PageError{"No sections found", nil}
	}

	return sections, nil
}

// GetLinks obtiene todos los enlaces de la página
func (h *SandboxPage) GetSandboxLinks() ([]string, error) {
	doc, err := h.fetchSandboxContent()
	if err != nil {
		return nil, err
	}

	links := make(map[string]bool)
	
	// Buscar enlaces en diferentes elementos
	selectors := []string{"a[href]", "link[href]", "[src]"}
	
	for _, selector := range selectors {
		doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
			if href, exists := s.Attr("href"); exists {
				links[href] = true
			}
			if src, exists := s.Attr("src"); exists {
				links[src] = true
			}
		})
	}

	// Convertir el mapa a slice
	uniqueLinks := make([]string, 0, len(links))
	for link := range links {
		uniqueLinks = append(uniqueLinks, link)
	}

	return uniqueLinks, nil
}

// VerifyStructure verifica la estructura de la página
func (h *SandboxPage) VerifySandboxStructure() (bool, error) {
	title, err := h.GetSandboxTitle()
	if err != nil {
		return false, err
	}

	if title != "Free Range Testers Sandbox" {
		return false, &PageError{
			Message: fmt.Sprintf("Unexpected title: %s", title),
			Err:     nil,
		}
	}

	sections, err := h.GetSandboxSections()
	if err != nil {
		return false, err
	}

	if len(sections) != 16 {
		return false, &PageError{
			Message: fmt.Sprintf("Expected 16 sections, found %d", len(sections)),
			Err:     nil,
		}
	}

	links, err := h.GetSandboxLinks()
	if err != nil {
		return false, err
	}

	if len(links) == 0 {
		return false, &PageError{
			Message: "No links found",
			Err:     nil,
		}
	}

	return true, nil
}

// ClickDynamicButton hace clic en un botón con ID dinámico
func (h *SandboxPage) ClickDynamicButton() (string,error) {
    // Crear un contexto de Chromedp
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

	// Crear un contexto con timeout para evitar bucles infinitos
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    

    // Navegar a la URL y hacer clic en el botón
    var popupVisible bool
    var hiddenText string

    err := chromedp.Run(ctx,
        chromedp.Navigate(h.URL),
		chromedp.WaitVisible(`button.btn.btn-primary`, chromedp.ByQuery),
        chromedp.Click(`button.btn.btn-primary`, chromedp.NodeVisible),		
		chromedp.WaitVisible(`#hidden-element`, chromedp.ByID),
        chromedp.Evaluate(`document.querySelector('#hidden-element') !== null`, &popupVisible),
        chromedp.Evaluate(`document.querySelector('#hidden-element').innerText`, &hiddenText),
    )
    if err != nil {
        return "",&PageError{"Error haciendo click al botón o esperando al texto oculto", err}
    }

    if !popupVisible {
        return "",&PageError{"No ha aparecido el texto oculto", nil}
    }

    return hiddenText, nil
}

// InsertTextInTextbox inserta texto en un cuadro de texto
func (h *SandboxPage) InsertTextInTextbox(text string) (string,error) {
    ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

    var insertedText string

	err := chromedp.Run(ctx,
		chromedp.Navigate(h.URL),
		chromedp.WaitVisible(`#formBasicText`, chromedp.ByID),
		chromedp.SetValue(`#formBasicText`, text, chromedp.ByID),
        chromedp.Value(`#formBasicText`, &insertedText, chromedp.ByID),        
	)
	if err != nil {
		return "",&PageError{"Error insertando texto en el cuadro de texto", err}
	}

    return insertedText, nil
}

// TestCheckboxesAndRadioButtons prueba los checkboxes y radio buttons
func (h *SandboxPage) TestCheckboxesAndRadioButtons() (string,string,error) {
    // Crear un contexto de Chromedp
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Crear un contexto con timeout para evitar bucles infinitos
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    var checkboxValue, radioValue string
    // Navegar a la URL y seleccionar los checkboxes y radio buttons
    err := chromedp.Run(ctx,
        chromedp.Navigate(h.URL),
        // Seleccionar los checkboxes
        chromedp.WaitVisible(`#checkbox-0`, chromedp.ByID),
        chromedp.Click(`#checkbox-0`, chromedp.NodeVisible),
        chromedp.Click(`#checkbox-1`, chromedp.NodeVisible),
        chromedp.Click(`#checkbox-2`, chromedp.NodeVisible),
        chromedp.Click(`#checkbox-3`, chromedp.NodeVisible),
        chromedp.Click(`#checkbox-4`, chromedp.NodeVisible),
        // Seleccionar los radio buttons
        chromedp.WaitVisible(`#formRadio1`, chromedp.ByID),
        chromedp.Click(`#formRadio1`, chromedp.NodeVisible),
        chromedp.WaitVisible(`#formRadio2`, chromedp.ByID),
        chromedp.Click(`#formRadio2`, chromedp.NodeVisible),
    )
    if err != nil {
        return "","",&PageError{"Error seleccionando checkboxes o radio buttons", err}
    }

    // Verificar que solo un radio button esté seleccionado
    var radio1Selected, radio2Selected bool
    err = chromedp.Run(ctx,
        chromedp.Evaluate(`document.querySelector('#formRadio1').checked`, &radio1Selected),
        chromedp.Evaluate(`document.querySelector('#formRadio2').checked`, &radio2Selected),
    )
    if err != nil {
        return "","",&PageError{"Error verificando selección de radio buttons", err}
    }

    if radio1Selected && radio2Selected {
        return "","",&PageError{"Ambos radio buttons están seleccionados, solo uno debería estarlo", nil}
    }
    // Obtener el valor del radio button seleccionado
    if radio1Selected {
        err = chromedp.Run(ctx,
            chromedp.Evaluate(`document.querySelector('#formRadio1').value`, &radioValue),
        )
    } else if radio2Selected {
        err = chromedp.Run(ctx,
            chromedp.Evaluate(`document.querySelector('#formRadio2').value`, &radioValue),
        )
    }
    if err != nil {
        return "", "", &PageError{"Error obteniendo el valor del radio button seleccionado", err}
    }

    // Obtener el valor de la etiqueta del primer checkbox seleccionado
    err = chromedp.Run(ctx,
        chromedp.Evaluate(`document.querySelector('#checkbox-0 + label').innerText`, &checkboxValue),
    )
    if err != nil {
        return "", "", &PageError{"Error obteniendo el valor de la etiqueta del checkbox seleccionado", err}
    }

    return checkboxValue, radioValue, nil
}

// ClickDropdowns hace clic en los dropdowns y selecciona opciones
func (h *SandboxPage) ClickDropdowns() (string, string, error) {
    // Crear un contexto de Chromedp
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Crear un contexto con timeout para evitar bucles infinitos
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    var firstDropdownValue, secondDropdownValue string
    // Navegar a la URL y seleccionar opciones en los dropdowns
    err := chromedp.Run(ctx,
        chromedp.Navigate(h.URL),
        // Seleccionar opción en el primer dropdown
        chromedp.WaitVisible(`#formBasicSelect`, chromedp.ByID),
        chromedp.SetValue(`#formBasicSelect`, "Fútbol", chromedp.ByID),
        // Cambia "Fútbol" por la opción que deseas seleccionar
        chromedp.Evaluate(`document.querySelector('#formBasicSelect').value`, &firstDropdownValue),
		// Seleccionar opción en el segundo dropdown
        chromedp.WaitVisible(`#dropdown-basic-button`, chromedp.ByID),
        chromedp.Click(`#dropdown-basic-button`, chromedp.NodeVisible),
        chromedp.WaitVisible(`.dropdown-menu`, chromedp.ByQuery),
        chromedp.Click(`.dropdown-menu a[href="#/action-2"]`, chromedp.NodeVisible), 
        // Cambia `href="#/action-2"` por l dia de la semana que deseas seleccionar
        chromedp.Evaluate(`document.querySelector('.dropdown-menu a[href="#/action-2"]').innerText`, &secondDropdownValue),
        //#root > div > div:nth-child(4) > div > div > div > a:nth-child(1)
        // Hacer clic en el botón de enviar
        chromedp.WaitVisible(`button.btn.btn-primary`, chromedp.ByQuery),
        chromedp.Click(`button.btn.btn-primary`, chromedp.NodeVisible),
    )
    if err != nil {
        return "", "", &PageError{"Error seleccionando opciones en los dropdowns", err}
    }

    return firstDropdownValue, secondDropdownValue, nil
}

// HandlePopup maneja el popup
func (h *SandboxPage) HandlePopup() (string, error) {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    var popupText string

    err := chromedp.Run(ctx,
        chromedp.Navigate(h.URL),
        //Esperamos a que cargue el Sandbox y pulsamos el botón de 'Mostrar Popup'
        chromedp.WaitVisible(`#root > div > div:nth-child(5) > div > button`, chromedp.ByQuery),
        chromedp.Click(`#root > div > div:nth-child(5) > div > button`, chromedp.NodeVisible),
        
        //Esperamos a que aparezca el popup. Guardamos el texto y pulsamos sobre 'Cerrar'
        chromedp.WaitVisible(`body > div.fade.modal.show > div > div`, chromedp.ByQuery),
        chromedp.Text(`body > div.fade.modal.show > div > div > div.modal-body`, &popupText, chromedp.BySearch), 
        chromedp.Click(`body > div.fade.modal.show > div > div > div.modal-footer > button`, chromedp.NodeVisible),
        
    )
    if err != nil {
        return "", &PageError{"Error manejando el popup", err}
    }
    //popupText = "popupText"
    return popupText, nil
}

// InteractWithShadowDOM interactúa con el Shadow DOM y devuelve su contenido
func (h *SandboxPage) InteractWithShadowDOM() (string, error) {
    // Crear un contexto de Chromedp
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Crear un contexto con timeout para evitar bucles infinitos
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    // Navegar a la URL y acceder al Shadow DOM
    var shadowContent string
    err := chromedp.Run(ctx,
        chromedp.Navigate(h.URL),
        chromedp.WaitVisible(`#shadow-root-example`, chromedp.ByID),
        chromedp.ActionFunc(func(ctx context.Context) error {
            // Acceder al shadow root
            var res string
            err := chromedp.Evaluate(`
                (function() {
                    const shadowHost = document.querySelector('#shadow-root-example');
                    if (!shadowHost) {
                        console.log('Shadow host no encontrado');
                        return '';
                    }
                    const shadowRoot = shadowHost.shadowRoot || shadowHost.attachShadow({mode: 'open'});
                    if (!shadowRoot) {
                        console.log('Shadow root no encontrado');
                        return '';
                    }
                    const shadowElement = shadowRoot.querySelector('#shadow-host');
                    if (!shadowElement) {
                        console.log('Elemento dentro del shadow root no encontrado');
                        return '';
                    }
                    return shadowElement.innerHTML;
                })()
            `, &res).Do(ctx)
            if err != nil {
                return err
            }
            shadowContent = res
            return nil
        }),
    )
    if err != nil {
        return "", &PageError{"Error interacting with Shadow DOM", err}
    }

    // Devolver el contenido del Shadow DOM
    return shadowContent, nil
}

// InteractWithTables interactúa con tablas dinámicas y estáticas
func (h *SandboxPage) InteractWithTables() (string, string, string, string, error) {
    // Crear un contexto de Chromedp
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Crear un contexto con timeout para evitar bucles infinitos
    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    var dynamicCellValueBefore, dynamicCellValueAfter, staticCellValueBefore, staticCellValueAfter string

    // Navegar a la URL y seleccionar opciones en los dropdowns
    err := chromedp.Run(ctx,
        chromedp.Navigate(h.URL),
        // Inspeccionamos la tabla dinámica
        chromedp.WaitVisible(`#root > div > div:nth-child(7) > div > table`, chromedp.BySearch),
        chromedp.Evaluate(`document.querySelector('#root > div > div:nth-child(7) > div > table').rows[1].cells[1].innerText`, &dynamicCellValueBefore),
        
        // Inspeccionamos la tabla estática
        chromedp.WaitVisible(`#root > div > div:nth-child(8) > div > table`, chromedp.BySearch),
        chromedp.Evaluate(`document.querySelector('#root > div > div:nth-child(8) > div > table').rows[1].cells[1].innerText`, &staticCellValueBefore),
    )
    if err != nil {
        return "", "", "","", &PageError{"Error inspeccionando las tablas", err}
    }

    // Recargar la página
    err = chromedp.Run(ctx,
        chromedp.Reload(),
        chromedp.WaitVisible(`#root > div > div:nth-child(7) > div > table`, chromedp.BySearch),
        chromedp.WaitVisible(`#root > div > div:nth-child(8) > div > table`, chromedp.BySearch),
        chromedp.Evaluate(`document.querySelector('#root > div > div:nth-child(7) > div > table').rows[1].cells[1].innerText`, &dynamicCellValueAfter),
        chromedp.Evaluate(`document.querySelector('#root > div > div:nth-child(8) > div > table').rows[1].cells[1].innerText`, &staticCellValueAfter),
    )   
    if err != nil {
        return "", "", "","", &PageError{"Error recargando la página y obteniendo el valor de la tabla dinámica", err}
    }

    return dynamicCellValueBefore, dynamicCellValueAfter, staticCellValueBefore, staticCellValueAfter, nil
}
