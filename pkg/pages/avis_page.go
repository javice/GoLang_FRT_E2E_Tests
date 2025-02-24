// pkg/pages/avis_page.go
package pages

import (
    "log"
    "time"

    "github.com/playwright-community/playwright-go"
)

// AvisPage representa la página de búsqueda de vehículos de Avis
type AvisPage struct {
    driver playwright.Page
}
func (p *AvisPage) Title() (string, error) {
    return p.driver.Title()
}

func (p *AvisPage) AvailableVehicles() ([]string, error) {
	VehicleList := p.driver.Locator(".vehicle")

	vehicles, err := VehicleList.AllTextContents()
	if err != nil {
		return nil, err
	}

	return vehicles, nil
	
}



// NewAvisPage crea una nueva instancia de AvisPage utilizando Playwright
func NewAvisPage() *AvisPage {
    pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("❌ Error al iniciar Playwright: %v", err)
    }

    browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: playwright.Bool(false),
    })
    if err != nil {
        log.Fatalf("❌ Error al abrir el navegador: %v", err)
    }

    page, err := browser.NewPage()
    if err != nil {
        log.Fatalf("❌ Error al crear la página: %v", err)
    }

    return &AvisPage{driver: page}
}

// Close cierra el navegador
func (ap *AvisPage) Close() {
    ap.driver.Close()
}

// NavigateTo navega a la URL especificada
func (ap *AvisPage) NavigateTo(url string) error {
    _, err := ap.driver.Goto(url)
    return err
}

// AcceptCookies acepta el popup emergente de cookies
func (ap *AvisPage) AcceptCookies() error {
    consentPromptLocator := ap.driver.Locator("#consent_prompt_accept")
    return consentPromptLocator.Click()
}

// SearchVehicles realiza la búsqueda de vehículos disponibles
func (ap *AvisPage) SearchVehicles(pickupTime, returnTime time.Time, pickupLocation, returnLocation string) error {
    if err := ap.selectPickupLocation(pickupLocation); err != nil {
        return err
    }

    if err := ap.selectReturnLocation(returnLocation); err != nil {
        return err
    }

    if err := ap.selectPickupDateTime(pickupTime); err != nil {
        return err
    }

    if err := ap.selectReturnDateTime(returnTime); err != nil {
        return err
    }

    return ap.simulateVehicleSearch()
}

func (ap *AvisPage) selectPickupLocation(pickupLocation string) error {
	pickupLocationLocator := ap.driver.Locator("#hire-search")
	if err := pickupLocationLocator.Click(); err != nil { return err }
	if err := pickupLocationLocator.PressSequentially(pickupLocation); err != nil {
		return err
	}

    time.Sleep(2 * time.Second)
	pickupSuggested := ap.driver.Locator("#getAQuote > div:nth-child(19) > div.standard-form__col.standard-form__col--init-full > div > ul > li:nth-child(1) > button")
	return pickupSuggested.Click()
}

func (ap *AvisPage) selectReturnLocation(returnLocation string) error {
	chkDifferentReturnLocation := ap.driver.Locator("#return-location-toggle > ul > li > label")
	if err := chkDifferentReturnLocation.Click(); err != nil {return err}
    
	returnLocationLocator := ap.driver.Locator("#return-search")
	if err := returnLocationLocator.Click(); err != nil { return err }
	if err := returnLocationLocator.PressSequentially(returnLocation); err != nil {
		return err
	}

    time.Sleep(2 * time.Second)
	returnSuggested := ap.driver.Locator("#getAQuote > div:nth-child(19) > div.standard-form__col.standard-form__col--init-hidden > div > ul > li:nth-child(1) > button")
    return returnSuggested.Click()
}


func (ap *AvisPage) selectPickupDateTime(pickupTime time.Time) error {
	pickupDateLocator := ap.driver.Locator("#date-from-display")
	if err := pickupDateLocator.Click(); err != nil {
		return err
	}
	pickupDateSuggested := ap.driver.Locator("#getAQuote > div.standard-form__row.booking-widget__date-fields > div:nth-child(1) > div > div.booking-widget__date-picker-container.booking-widget__date-picker-container--open > div > div > div:nth-child(1) > table > tbody > tr:nth-child(5) > td.is-selected > button")
	if err := pickupDateSuggested.Click(); err != nil {
		return err
	}
	pickupTimeLocator := ap.driver.Locator("#time-from-display")
	if err := pickupTimeLocator.Click(); err != nil {
		return err
	}
	pickupTimeSuggested := ap.driver.Locator("#getAQuote > div.standard-form__row.booking-widget__date-fields > div:nth-child(1) > div > div.booking-widget__time-picker-container > div > div > ul > li.ui-timepicker-am.ui-timepicker-selected")
	return pickupTimeSuggested.Click()
}

func (ap *AvisPage) selectReturnDateTime(returnTime time.Time) error {
	returnDateLocator := ap.driver.Locator("#date-to-display")
	if err := returnDateLocator.Click(); err != nil {
		return err
	}
	returnDateSuggested := ap.driver.Locator("#getAQuote > div.standard-form__row.booking-widget__date-fields > div:nth-child(2) > div > div.booking-widget__date-picker-container.booking-widget__date-picker-container--open > div > div > div:nth-child(1) > table > tbody > tr:nth-child(5) > td.is-selected > button")
	if err := returnDateSuggested.Click(); err != nil {
		return err
	}
	returnTimeLocator := ap.driver.Locator("#time-to-display")
	if err := returnTimeLocator.Click(); err != nil {
		return err
	}
	returnTimeSuggested := ap.driver.Locator("#getAQuote > div.standard-form__row.booking-widget__date-fields > div:nth-child(2) > div > div.booking-widget__time-picker-container > div > div > ul > li.ui-timepicker-am.ui-timepicker-selected")
	return returnTimeSuggested.Click()
}



func (ap *AvisPage) simulateVehicleSearch() error {
    // Asegurarse de que el botón esté visible y habilitado antes de hacer clic
    btnBuscar := ap.driver.GetByRole("button", playwright.PageGetByRoleOptions{
        Name: "Buscar",
    })
    if err := btnBuscar.WaitFor(playwright.LocatorWaitForOptions{
        State:   playwright.WaitForSelectorStateVisible,
        Timeout: playwright.Float(20000), // 20 segundos de tiempo de espera
    }); err != nil {
        return err
    }
    if err := btnBuscar.Click(); err != nil {
        return err
    }
    titleLocator := ap.driver.Locator("#title__heading")
    err := titleLocator.WaitFor(playwright.LocatorWaitForOptions{
        State:   playwright.WaitForSelectorStateVisible,
        Timeout: playwright.Float(20000), // 20 segundos de tiempo de espera
    })
    return err
}

