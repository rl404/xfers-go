package main

import (
	"log"

	"github.com/rl404/xfers-go"
)

func main() {
	apiKey := "test_xxx"
	secretKey := "abc123"

	x := xfers.NewDefault(apiKey, secretKey, xfers.Sandbox)

	bank, code, err := x.ValidateBankAccount(xfers.ValidateBankAccountRequest{
		AccountNo:     "000501003219303",
		BankShortCode: xfers.BankBRI,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, bank)
}
