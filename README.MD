# GoLang_FRT_E2E_Tests

Este proyecto es un conjunto de pruebas de extremo a extremo (E2E) para FreeRangeTesters, escritas en Go.

## Estructura del proyecto

La estructura del proyecto es la siguiente:

```mermaid
graph TD;

    root[.]
    root --> cmd
    root --> pkg
    root --> tests
    root --> reports

.
├── cmd
│   ├── generate_report
│   │   └── main.go
│   └── run_tests
│       └── main.go
├── pkg
│   ├── pages
│   │   ├── home_page.go
│   │   ├── sandbox_page.go
│   │   └── ...
│   └── reports
│       ├── test_report.go
│       └── ...
├── tests
│   ├── e2e
│   │   ├── home_page_test.go
│   │   ├── sandbox_page_test.go
│   │   └── ...
│   └── ...
├── reports
│   ├── test-report.html
│   └── test.log
|   └── ...    
├── Makefile
├── go.mod
├── go.sum
└── README.
```

* `cmd/generate_report`: contiene el comando para generar informes de pruebas.
* `Makefile`: contiene las reglas de construcción para el proyecto.
* `cmd`: contiene los comandos para ejecutar las pruebas y generar informes.
* `pkg`: contiene los paquetes de Go que se utilizan en el proyecto.
  * `pages`: contiene las definiciones de las páginas que se prueban. Siguiendo elo modelo POM (Page Object Model).
  * `reports`: contiene las funciones para generar informes de pruebas.
* `tests`: contiene las pruebas de extremo a extremo.
  * `e2e`: contiene las pruebas de extremo a extremo para las páginas de FreeRangeTesters.
* `reports`: contiene los informes de pruebas generados.

## Requisitos previos

* Go instalado en el sistema.
* Dependencias de Go instaladas (`go mod download`).

## Ejecución de las pruebas

1. Ejecuta `go test -v ./...` para ejecutar todas las pruebas.
2. Ejecuta `go run cmd/generate_report/main.go` para generar un informe de pruebas.
3. Ejecuta 'make test-report' para ejecutar todas las pruebas y generar un informe de pruebas en formato HTML.

## Informes de pruebas

Los informes de pruebas se generan en el directorio `reports`. El archivo `test-report.html` contiene el informe de pruebas en formato HTML.

## Contribución

Si deseas contribuir al proyecto, por favor, crea un fork del repositorio y envía una solicitud de pull request con tus cambios o sugerencias.
