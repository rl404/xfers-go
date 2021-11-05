# Xfers-go

[![Go Report Card](https://goreportcard.com/badge/github.com/rl404/xfers-go)](https://goreportcard.com/report/github.com/rl404/xfers-go)
![License: MIT](https://img.shields.io/github/license/rl404/xfers-go.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/rl404/xfers-go.svg)](https://pkg.go.dev/github.com/rl404/xfers-go)

_xfers-go_ is unofficial golang API wrapper for [xfers](https://www.xfers.com/id?). Only for V4 indonesian xfers.

For official documentation, go [here](https://docs.xfers.com/reference/introduction-1).

## Features

- Get xfers account's balance
- Get disbursement bank list
- Validate bank account
- Create payment
  - virtual account
  - retail outlet
  - QRIS
  - e-wallet
- Get payment
- Get payment list + filter + pagination
- Simulate payment
- Create disbursement
- Get disbursement
- Get disbursement list + filter + pagination

## Installation

```
go get github.com/rl404/xfers-go
```

## Quick Start

```go
package main

import (
	"log"

	"github.com/rl404/xfers-go"
)

func main() {
    // Prepare API and secret key.
	apiKey := "test_xxx"
	secretKey := "abc123"

    // Create xfers client.
	x := xfers.NewDefault(apiKey, secretKey, xfers.Sandbox)

    // Get your xfers's account balance.
	balance, code, err := x.GetBalance()
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, balance)
}
```

*For more detail config and usage, please go to the [documentation](https://pkg.go.dev/github.com/rl404/xfers-go).*

## License

MIT License

Copyright (c) 2021 Axel
