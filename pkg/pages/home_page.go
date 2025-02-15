// pkg/pages/home_page.go
package pages

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// PageError representa los errores personalizados para el scraping
type PageError struct {
	Message string
	Err     error
}

func (e *PageError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// HomePage representa la página principal
type HomePage struct {
	URL    string
	client *http.Client
}

// NewHomePage crea una nueva instancia de HomePage
func NewHomePage() *HomePage {
	return &HomePage{
		URL: "https://www.freerangetesters.com",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// fetchContent obtiene el contenido de la página
func (h *HomePage) fetchContent() (*goquery.Document, error) {
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
func (h *HomePage) GetTitle() (string, error) {
	doc, err := h.fetchContent()
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
func (h *HomePage) GetSections() ([]string, error) {
	doc, err := h.fetchContent()
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
func (h *HomePage) GetLinks() ([]string, error) {
	doc, err := h.fetchContent()
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
func (h *HomePage) VerifyStructure() (bool, error) {
	title, err := h.GetTitle()
	if err != nil {
		return false, err
	}

	if title != "Free Range Testers" {
		return false, &PageError{
			Message: fmt.Sprintf("Unexpected title: %s", title),
			Err:     nil,
		}
	}

	sections, err := h.GetSections()
	if err != nil {
		return false, err
	}

	if len(sections) != 16 {
		return false, &PageError{
			Message: fmt.Sprintf("Expected 16 sections, found %d", len(sections)),
			Err:     nil,
		}
	}

	links, err := h.GetLinks()
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