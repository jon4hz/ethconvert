[![Go Report Card](https://goreportcard.com/badge/github.com/jon4hz/ethconvert)](https://goreportcard.com/report/github.com/jon4hz/ethconvert)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
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

## CLI
### Install
`go install github.com/jon4hz/ethconvert/cmd/ethconvert@latest`

### Example
```bash
$ ethconvert -i ether -o wei -a 1
1000000000000000000 wei

$ ethconvert 50 wei gwei
0.00000005 gwei
```