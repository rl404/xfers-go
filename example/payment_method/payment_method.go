package main

import (
	"log"

	"github.com/rl404/xfers-go"
)

func main() {
	apiKey := "test_xxx"
	secretKey := "abc123"

	x := xfers.NewDefault(apiKey, secretKey, xfers.Sandbox)

	payment, code, err := x.CreatePaymentMethod(xfers.CreatePaymentMethodRequest{
		Type:          xfers.PaymentVA,
		ReferenceID:   "uuid-payment2-1",
		DisplayName:   "display name",
		BankShortCode: xfers.BankMandiri,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, payment)

	payment, code, err = x.GetPaymentMethod(xfers.GetPaymentMethodRequest{
		Type: xfers.PaymentVA,
		ID:   payment.ID,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, payment)

	payments, code, err := x.GetPaymentMethods(
		xfers.GetPaymentMethodRequest{
			Type: xfers.PaymentVA,
			ID:   "va_5fe9457c78aeadc7ede490acd26aea54",
		},
		xfers.Pagination{
			Page:     1,
			PageSize: 2,
			Sort:     "createdAt",
			// CreatedAfter:  time.Now().Add(-1 * time.Hour),
			// CreatedBefore: time.Now().Add(time.Hour),
			// Status:        xfers.StatusProcessing,
			// ReferenceID:   "uuid-payment-6",
		})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, payments)

	action, code, err := x.SimulatePaymentMethod(xfers.SimulatePaymentMethodRequest{
		ID:     payment.ID,
		Type:   xfers.PaymentVA,
		Action: xfers.ActionReceivePayment,
		Amount: 10000,
	})
	if err != nil {
		log.Println(code, err)
		return
	}

	log.Println(code, action)
}
