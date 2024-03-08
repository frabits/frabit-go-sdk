# frabit-go-sdk
[![GitHub release](https://img.shields.io/github/v/release/frabits/frabit-go-sdk)](https://github.com/frabits/frabit-go-sdk/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/frabits/frabit-go-sdk?utm_source=godoc)](https://godoc.org/github.com/frabits/frabit-go-sdk)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/frabits/frabit-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/frabits/frabit-go-sdk)](https://goreportcard.com/report/github.com/frabits/frabit-go-sdk)
![GitHub](https://img.shields.io/github/license/frabits/frabit-go-sdk)


Frabit official golang sdk

# Installation
```bash
go get https://github.com/frabits/frabit-go-sdk
```

# Examples

```golang
package main

import (
	"os"
	
	fb "github.com/frabits/frabit-go-sdk"
)

func main() {
	baseUrl := os.Getenv("FRABIT_BASE_URL")
	token := os.Getenv("FRABIT_TOKEN")
	
	client := fb.NewClient(fb.WithBaseUrl(baseUrl), fb.WithToken(token))

	client.database.get()
}
```