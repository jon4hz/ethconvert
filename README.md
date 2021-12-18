# ethconvert

Ethereum unit converter written in go.

## Install

`go get github.com/jon4hz/ethconvert/pkg/ethconvert`

## Example usage

```go
package main

import (
	"fmt"

	"github.com/jon4hz/ethconvert/pkg/ethconvert"
	"github.com/shopspring/decimal"
)

func main() {
	amount := decimal.NewFromInt(50)
	gwei, err := ethconvert.Convert(amount, "ether", "gwei")
	if err != nil {
		panic(err)
	}
	fmt.Println(gwei)
}
```

