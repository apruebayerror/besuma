# Besuma

Besuma prueba de una dependencia GO que suma.

## Installation

En la terminal del proyecto ejecutar el siguiente comando.

```bash
go get github.com/apruebayerror/besuma
```

## Usage

```go
package main

import (
	"fmt"

	sum "github.com/apruebayerror/besuma"
)

func main() {
	valor := sum.Suma(10)
	fmt.Println(valor)
}
```
