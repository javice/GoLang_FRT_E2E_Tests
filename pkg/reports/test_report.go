// test_report.go

// Package reports contiene funciones para generar reportes de pruebas E2E en FreeRangeTesters

package reports

import (
    "html/template"
    "os"
    "strings"
    "time"
)

type TestResult struct {
    Name      string
    Status    string
    Logs      []string
    Timestamp time.Time
    Duration  time.Duration
    SubTests  []*TestResult
}

func lower(s string) string {
    return strings.ToLower(s)
}

func GenerateHTMLReport(results []TestResult, outputPath string) error {
    const tpl = `
    <!DOCTYPE html>
    <html lang="es">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Reporte HTML de Pruebas E2E en FreeRangeTesters</title>
        <style>
            body { font-family: Arial, sans-serif; margin: 20px; }
            .test { margin: 20px 0; padding: 15px; border: 1px solid #ddd; border-radius: 5px; }
            .pass { background-color: #e8f5e9; }
            .fail { background-color: #ffebee; }
            .log { background-color: #f5f5f5; padding: 10px; margin-top: 10px; font-family: monospace; }
            .timestamp { color: #666; font-size: 0.9em; }
            .duration { color: #666; font-style: italic; }
            .error { color: red; margin-top: 10px; }
            .subtest { margin-left: 20px; }
            .running { color: yellow; }
        </style>
    </head>
    <body>
        <h1>Reporte HTML de Pruebas E2E en FreeRangeTesters</h1>
        {{range .}}
        <div class="test">
            <h2>{{.Name}} - <span class="{{.Status | lower}}">{{.Status}}</span></h2>
            <p class="timestamp">Inicio: {{.Timestamp.Format "2006-01-02 15:04:05"}}</p>
            <p class="duration">Duración: {{.Duration.Seconds}}s</p>
            <div class="log">
                {{range .Logs}}
                {{.}}<br>
                {{end}}
            </div>
            {{if .SubTests}}
            <div class="subtest">
                {{range .SubTests}}
                <h3>{{.Name}} - <span class="{{.Status | lower}}">{{.Status}}</span></h3>
                <p class="duration">Duración: {{.Duration.Seconds}}s</p>
                <div class="log">
                    {{range .Logs}}
                    {{.}}<br>
                    {{end}}
                </div>
                {{end}}
            </div>
            {{end}}
        </div>
        {{end}}
    </body>
    </html>
    `

    tmpl, err := template.New("report").Funcs(template.FuncMap{"lower": lower}).Parse(tpl)
    if err != nil {
        return err
    }

    file, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer file.Close()

    return tmpl.Execute(file, results)
}
