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

	payment, code, err := x.CreatePayment(xfers.CreatePaymentRequest{
		PaymentMethodType: xfers.PaymentVA,
		Amount:            20000,
		ReferenceID:       "uuid-payment-7",
		ExpiredAt:         time.Now().Add(2 * time.Hour),
		Description:       "description",
		DisplayName:       "display name",
		// RetailOutletName:  xfers.OutletAlfamart,
		BankShortCode: xfers.BankBNI,
		SuffixNo:      "123",
		// ProvideCode:              xfers.EWalletShopeePay,
		// AfterSettlementReturnURL: "http://google.com",
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, payment)

	payment, code, err = x.GetPayment(payment.ID)
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, payment)

	payments, code, err := x.GetPayments(xfers.Pagination{
		Page:          1,
		PageSize:      2,
		Sort:          "createdAt",
		CreatedAfter:  time.Now().Add(-1 * time.Hour),
		CreatedBefore: time.Now().Add(time.Hour),
		Status:        xfers.StatusProcessing,
		ReferenceID:   "uuid-payment-6",
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, payments)

	action, code, err := x.SimulatePayment(xfers.SimulatePaymentRequest{
		ID:     payment.ID,
		Action: xfers.ActionReceivePayment,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, action)
}
