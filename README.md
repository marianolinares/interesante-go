## Prerequisites

You'll need GoLang `1.20` installed.
Install azure functions.

## Getting Started

**Add initial modules**

```
go mod tidy
```

**Start Service**

```
cd code/cmd/functions
go build main.go
func start
```

## Hexagonal Architecture

* code/internal: domain.
* first level folders (code/internal/** except platform): application.
* second level folders (in code/internal/platform): infrastructure.
