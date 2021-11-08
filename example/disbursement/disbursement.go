package main

import (
	"log"
	"time"

	"github.com/rl404/xfers-go"
)

func main() {
	apiKey := "test_xxx"
	secretKey := "abc123"

	x := xfers.NewDefault(apiKey, secretKey, xfers.Sandbox)

	disbursement, code, err := x.CreateDisbursement(xfers.CreateDisbursementRequest{
		ReferenceID:           "uuid-8",
		BankAccountHolderName: "Name",
		BankAccountNo:         "123",
		BankShortCode:         xfers.BankBCA,
		Amount:                10000,
		Description:           "desc",
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, disbursement)

	disbursement, code, err = x.GetDisbursement(disbursement.ID)
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, disbursement)

	disbursements, code, err := x.GetDisbursements(xfers.Pagination{
		Page:          1,
		PageSize:      2,
		Sort:          "createdAt",
		CreatedAfter:  time.Now().Add(-1 * time.Hour),
		CreatedBefore: time.Now().Add(time.Hour),
		Status:        xfers.StatusProcessing,
		ReferenceID:   "uuid-2",
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, disbursements)

	action, code, err := x.SimulateDisbursement(xfers.SimulateDisbursementRequest{
		ID:     disbursement.ID,
		Action: xfers.ActionFail,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, action)
}
