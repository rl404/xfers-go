package xfers

import (
	"context"
	"fmt"
	"net/http"
)

// PaymentMethod is payment method model.
type PaymentMethod struct {
	ID            string
	Type          PaymentType
	ReferenceID   string
	DisplayName   string
	BankShortCode BankCode // va
	AccountNo     string   // va
	ImageURL      string   // qris
}

// CreatePaymentMethod to create new payment method.
func (c *Client) CreatePaymentMethod(request CreatePaymentMethodRequest) (*PaymentMethod, int, error) {
	return c.CreatePaymentMethodWithContext(context.Background(), request)
}

// CreatePaymentMethodWithContext to create new payment with context.
func (c *Client) CreatePaymentMethodWithContext(ctx context.Context, request CreatePaymentMethodRequest) (*PaymentMethod, int, error) {
	if err := request.validate(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response paymentMethod
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/payment_methods/%s", c.baseURL, request.Type.toURL()),
		c.apiKey,
		c.secretKey,
		nil,
		request.wrap(),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toPaymentMethod(), code, nil
}

// GetPaymentMethod to get payment method.
func (c *Client) GetPaymentMethod(request GetPaymentMethodRequest) (*PaymentMethod, int, error) {
	return c.GetPaymentMethodWithContext(context.Background(), request)
}

// GetPaymentMethodWithContext to get payment with context.
func (c *Client) GetPaymentMethodWithContext(ctx context.Context, request GetPaymentMethodRequest) (*PaymentMethod, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response paymentMethod
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/payment_methods/%s/%s", c.baseURL, request.Type.toURL(), request.ID),
		c.apiKey,
		c.secretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toPaymentMethod(), code, nil
}

// GetPaymentMethods to get payment method list.
func (c *Client) GetPaymentMethods(request GetPaymentMethodRequest, pagination Pagination) ([]Payment, int, error) {
	return c.GetPaymentMethodsWithContext(context.Background(), request, pagination)
}

// GetPaymentMethodsWithContext to get payment method list with context.
func (c *Client) GetPaymentMethodsWithContext(ctx context.Context, request GetPaymentMethodRequest, pagination Pagination) ([]Payment, int, error) {
	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	if err := validate(&pagination); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response payments
	code, err := c.requester.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/payment_methods/%s/%s/payments?%s", c.baseURL, request.Type.toURL(), request.ID, pagination.encode()),
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

// PaymentMethodAction is response model from simulate payment method.
type PaymentMethodAction struct {
	TargetID   string
	TargetType string
	Name       Action
}

// SimulatePaymentMethod to simulate payment method status. Sandbox only.
func (c *Client) SimulatePaymentMethod(request SimulatePaymentMethodRequest) (*PaymentMethodAction, int, error) {
	return c.SimulatePaymentWithMethodContext(context.Background(), request)
}

// SimulatePaymentWithMethodContext to simulate payment method status with context. Sandbox only.
func (c *Client) SimulatePaymentWithMethodContext(ctx context.Context, request SimulatePaymentMethodRequest) (*PaymentMethodAction, int, error) {
	if c.env == Production {
		return nil, http.StatusBadRequest, ErrSandboxOnly
	}

	if err := validate(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	var response paymentMethodAction
	code, err := c.requester.Call(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/payment_methods/%s/%s/tasks", c.baseURL, request.Type.toURL(), request.ID),
		c.apiKey,
		c.secretKey,
		nil,
		request.wrap(),
		&response,
	)
	if err != nil {
		return nil, code, err
	}

	return response.toPaymentMethodAction(), code, nil
}
