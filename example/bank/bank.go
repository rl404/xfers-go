package main

import (
	"log"

	"github.com/rl404/xfers-go"
)

func main() {
	apiKey := "test_xxx"
	secretKey := "abc123"

	x := xfers.NewDefault(apiKey, secretKey, xfers.Sandbox)

	banks, code, err := x.GetBanks()
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, banks)
}
