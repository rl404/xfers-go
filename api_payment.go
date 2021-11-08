package xfers

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Payment is payment model.
type Payment struct {
	ID                 string
	ReferenceID        string
	PaymentMethodID    string
	Type               PaymentType
	Amount             float64
	Fees               float64
	Status             Status
	Description        string
	DisplayName        string
	RetailOutletCode   RetailOutlet // retail
	PaymentCode        string       // retail
	BankShortCode      BankCode     // va
	AccountNo          string       // va
	ImageURL           string       // qris
	HttpURL            string       // e-wallet
	AfterSettlementURl string       // e-wallet
	ExpiredAt          time.Time
	CreatedAt          time.Time
}

// CreatePayment to create new payment.
func (c *Client) CreatePayment(request CreatePaymentRequest) (*Payment, int, error) {
	return c.CreatePaymentWithContext(context.Background(), request)
}

// CreatePaymentWithContext to create new payment with context.
func (c *Client) CreatePaymentWithContext(ctx context.Context, request CreatePaymentRequest) (*Payment, int, error) {
	if err := request.validate(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response payment
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/payments", c.baseURL),
		c.apiKey,
		c.secretKey,
		nil,
		request.wrap(),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toPayment(), code, nil
}

// GetPayment to get payment.
func (c *Client) GetPayment(id string) (*Payment, int, error) {
	return c.GetPaymentWithContext(context.Background(), id)
}

// GetPaymentWithContext to get payment with context.
func (c *Client) GetPaymentWithContext(ctx context.Context, id string) (*Payment, int, error) {
	if id == "" {
		return nil, http.StatusBadRequest, errRequiredField("id")
	}

	var response payment
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/payments/%s", c.baseURL, id),
		c.apiKey,
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toPayment(), code, nil
}

// GetPayments to get payment list.
func (c *Client) GetPayments(request Pagination) ([]Payment, int, error) {
	return c.GetPaymentsWithContext(context.Background(), request)
}

// GetPaymentsWithContext to get disbursement list with context.
func (c *Client) GetPaymentsWithContext(ctx context.Context, request Pagination) ([]Payment, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response payments
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/payments?%s", c.baseURL, request.encode()),
		c.apiKey,
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toPayments(), code, nil
}

// PaymentAction is response model from simulate payment.
type PaymentAction struct {
	TargetID   string
	TargetType string
	Action     Action
}

// SimulatePayment to simulate payment status. Sandbox only.
func (c *Client) SimulatePayment(request SimulatePaymentRequest) (*PaymentAction, int, error) {
	return c.SimulatePaymentWithContext(context.Background(), request)
}

// SimulatePaymentWithContext to simulate payment status with context. Sandbox only.
func (c *Client) SimulatePaymentWithContext(ctx context.Context, request SimulatePaymentRequest) (*PaymentAction, int, error) {
	if c.env == Production {
		return nil, http.StatusBadRequest, ErrSandboxOnly
	}

	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response paymentAction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/payments/%s/tasks", c.baseURL, request.ID),
		c.apiKey,
		c.secretKey,
		nil,
		request.wrap(),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toPaymentAction(), code, nil
}
