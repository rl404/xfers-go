package xfers

import (
	"strconv"
	"time"
)

type balance struct {
	Data struct {
		Attributes struct {
			TotalBalance     string `json:"totalBalance"`
			AvailableBalance string `json:"availableBalance"`
			PendingBalance   string `json:"pendingBalance"`
		} `json:"attributes"`
	} `json:"data"`
}

func (b balance) toBalance() *Balance {
	tb, _ := strconv.ParseFloat(b.Data.Attributes.TotalBalance, 64)
	ab, _ := strconv.ParseFloat(b.Data.Attributes.AvailableBalance, 64)
	pb, _ := strconv.ParseFloat(b.Data.Attributes.PendingBalance, 64)
	return &Balance{
		TotalBalance:     tb,
		AvailableBalance: ab,
		PendingBalance:   pb,
	}
}

type bank struct {
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name      string   `json:"name"`
			ShortCode BankCode `json:"shortCode"`
		} `json:"attributes"`
	} `json:"data"`
}

func (b *bank) toBanks() []Bank {
	banks := make([]Bank, len(b.Data))
	for i, bb := range b.Data {
		banks[i] = Bank{
			Name:      bb.Attributes.Name,
			ShortCode: bb.Attributes.ShortCode,
		}
	}
	return banks
}

type bankAccount struct {
	Data struct {
		Attributes struct {
			AccountName   string   `json:"accountName"`
			AccountNo     string   `json:"accountNo"`
			BankShortCode BankCode `json:"bankShortCode"`
		} `json:"attributes"`
	} `json:"data"`
}

func (b bankAccount) toBankAccount() *BankAccount {
	return &BankAccount{
		AccountName:   b.Data.Attributes.AccountName,
		AccountNo:     b.Data.Attributes.AccountNo,
		BankShortCode: b.Data.Attributes.BankShortCode,
	}
}

type disbursement struct {
	Data disbursementData `json:"data"`
}

type disbursementData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		ReferenceID        string    `json:"referenceId"`
		Description        string    `json:"description"`
		Amount             string    `json:"amount"`
		Status             Status    `json:"status"`
		CreatedAt          time.Time `json:"createdAt"`
		Fees               string    `json:"fees"`
		FailureReason      string    `json:"failureReason"`
		DisbursementMethod struct {
			Type                        DisbursementType `json:"type"`
			BankAccountNo               string           `json:"bankAccountNo"`
			BankShortCode               BankCode         `json:"bankShortCode"`
			BankName                    string           `json:"bankName"`
			BankAccountHolderName       string           `json:"bankAccountHolderName"`
			ServerBankAccountHolderName string           `json:"serverBankAccountHolderName"`
		} `json:"disbursementMethod"`
	} `json:"attributes"`
}

func (d disbursement) toDisbursement() *Disbursement {
	amount, _ := strconv.ParseFloat(d.Data.Attributes.Amount, 64)
	fees, _ := strconv.ParseFloat(d.Data.Attributes.Fees, 64)
	return &Disbursement{
		ID:                          d.Data.ID,
		ReferenceID:                 d.Data.Attributes.ReferenceID,
		Description:                 d.Data.Attributes.Description,
		Amount:                      amount,
		Status:                      d.Data.Attributes.Status,
		CreatedAt:                   d.Data.Attributes.CreatedAt,
		Fees:                        fees,
		Type:                        d.Data.Attributes.DisbursementMethod.Type,
		FailureReason:               d.Data.Attributes.FailureReason,
		BankAccountNo:               d.Data.Attributes.DisbursementMethod.BankAccountNo,
		BankShortCode:               d.Data.Attributes.DisbursementMethod.BankShortCode,
		BankName:                    d.Data.Attributes.DisbursementMethod.BankName,
		BankAccountHolderName:       d.Data.Attributes.DisbursementMethod.BankAccountHolderName,
		ServerBankAccountHolderName: d.Data.Attributes.DisbursementMethod.ServerBankAccountHolderName,
	}
}

type disbursements struct {
	Data []disbursementData `json:"data"`
}

func (d disbursements) toDisbursements() []Disbursement {
	disbursements := make([]Disbursement, len(d.Data))
	for i, dd := range d.Data {
		amount, _ := strconv.ParseFloat(dd.Attributes.Amount, 64)
		fees, _ := strconv.ParseFloat(dd.Attributes.Fees, 64)
		disbursements[i] = Disbursement{
			ID:                          dd.ID,
			ReferenceID:                 dd.Attributes.ReferenceID,
			Description:                 dd.Attributes.Description,
			Amount:                      amount,
			Status:                      dd.Attributes.Status,
			CreatedAt:                   dd.Attributes.CreatedAt,
			Fees:                        fees,
			Type:                        dd.Attributes.DisbursementMethod.Type,
			FailureReason:               dd.Attributes.FailureReason,
			BankAccountNo:               dd.Attributes.DisbursementMethod.BankAccountNo,
			BankShortCode:               dd.Attributes.DisbursementMethod.BankShortCode,
			BankName:                    dd.Attributes.DisbursementMethod.BankName,
			BankAccountHolderName:       dd.Attributes.DisbursementMethod.BankAccountHolderName,
			ServerBankAccountHolderName: dd.Attributes.DisbursementMethod.ServerBankAccountHolderName,
		}
	}
	return disbursements
}

type disbursementAction struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			TargetID   string `json:"targetId"`
			TargetType string `json:"targetType"`
			Action     Action `json:"action"`
		} `json:"attributes"`
	} `json:"data"`
}

func (d disbursementAction) toDisbursementAction() *DisbursementAction {
	return &DisbursementAction{
		TargetID:   d.Data.Attributes.TargetID,
		TargetType: d.Data.Attributes.TargetType,
		Action:     d.Data.Attributes.Action,
	}
}

type payment struct {
	Data paymentData `json:"data"`
}

type paymentData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Status        Status    `json:"status"`
		Amount        string    `json:"amount"`
		CreatedAt     time.Time `json:"createdAt"`
		Description   string    `json:"description"`
		ExpiredAt     time.Time `json:"expiredAt"`
		ReferenceID   string    `json:"referenceId"`
		Fees          string    `json:"fees"`
		PaymentMethod struct {
			ID           string      `json:"id"`
			Type         PaymentType `json:"type"`
			ReferenceID  string      `json:"referenceId"`
			Instructions struct {
				DisplayName string `json:"displayName"`

				// Retail outlet.
				RetailOutletCode RetailOutlet `json:"retailOutletCode"`
				PaymentCode      string       `json:"paymentCode"`

				// VA.
				BankShortCode BankCode `json:"bankShortCode"`
				AccountNo     string   `json:"accountNo"`

				// QRIS.
				ImageURL string `json:"imageUrl"`
			} `json:"instructions"`
			// E-wallet.
			Settlement struct {
				HttpURL            string `json:"httpUrl"`
				AfterSettlementURL string `json:"afterSettlementUrl"`
			} `json:"settlement"`
		} `json:"paymentMethod"`
	} `json:"attributes"`
}

func (p payment) toPayment() *Payment {
	amount, _ := strconv.ParseFloat(p.Data.Attributes.Amount, 64)
	fees, _ := strconv.ParseFloat(p.Data.Attributes.Fees, 64)
	return &Payment{
		ID:                 p.Data.ID,
		Status:             p.Data.Attributes.Status,
		Amount:             amount,
		CreatedAt:          p.Data.Attributes.CreatedAt,
		Description:        p.Data.Attributes.Description,
		ExpiredAt:          p.Data.Attributes.ExpiredAt,
		ReferenceID:        p.Data.Attributes.ReferenceID,
		Fees:               fees,
		PaymentMethodID:    p.Data.Attributes.PaymentMethod.ID,
		Type:               p.Data.Attributes.PaymentMethod.Type,
		DisplayName:        p.Data.Attributes.PaymentMethod.Instructions.DisplayName,
		RetailOutletCode:   p.Data.Attributes.PaymentMethod.Instructions.RetailOutletCode,
		PaymentCode:        p.Data.Attributes.PaymentMethod.Instructions.PaymentCode,
		BankShortCode:      p.Data.Attributes.PaymentMethod.Instructions.BankShortCode,
		AccountNo:          p.Data.Attributes.PaymentMethod.Instructions.AccountNo,
		ImageURL:           p.Data.Attributes.PaymentMethod.Instructions.ImageURL,
		HttpURL:            p.Data.Attributes.PaymentMethod.Settlement.HttpURL,
		AfterSettlementURl: p.Data.Attributes.PaymentMethod.Settlement.AfterSettlementURL,
	}
}

type payments struct {
	Data []paymentData `json:"data"`
}

func (p payments) toPayments() []Payment {
	payments := make([]Payment, len(p.Data))
	for i, pp := range p.Data {
		amount, _ := strconv.ParseFloat(pp.Attributes.Amount, 64)
		fees, _ := strconv.ParseFloat(pp.Attributes.Fees, 64)
		payments[i] = Payment{
			ID:                 pp.ID,
			Status:             pp.Attributes.Status,
			Amount:             amount,
			CreatedAt:          pp.Attributes.CreatedAt,
			Description:        pp.Attributes.Description,
			ExpiredAt:          pp.Attributes.ExpiredAt,
			ReferenceID:        pp.Attributes.ReferenceID,
			Fees:               fees,
			PaymentMethodID:    pp.Attributes.PaymentMethod.ID,
			Type:               pp.Attributes.PaymentMethod.Type,
			DisplayName:        pp.Attributes.PaymentMethod.Instructions.DisplayName,
			RetailOutletCode:   pp.Attributes.PaymentMethod.Instructions.RetailOutletCode,
			PaymentCode:        pp.Attributes.PaymentMethod.Instructions.PaymentCode,
			BankShortCode:      pp.Attributes.PaymentMethod.Instructions.BankShortCode,
			AccountNo:          pp.Attributes.PaymentMethod.Instructions.AccountNo,
			ImageURL:           pp.Attributes.PaymentMethod.Instructions.ImageURL,
			HttpURL:            pp.Attributes.PaymentMethod.Settlement.HttpURL,
			AfterSettlementURl: pp.Attributes.PaymentMethod.Settlement.AfterSettlementURL,
		}
	}
	return payments
}

type paymentAction struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			TargetID   string `json:"targetId"`
			TargetType string `json:"targetType"`
			Action     Action `json:"action"`
		} `json:"attributes"`
	} `json:"data"`
}

func (p paymentAction) toPaymentAction() *PaymentAction {
	return &PaymentAction{
		TargetID:   p.Data.Attributes.TargetID,
		TargetType: p.Data.Attributes.TargetType,
		Action:     p.Data.Attributes.Action,
	}
}

type paymentMethod struct {
	Data struct {
		ID         string      `json:"id"`
		Type       PaymentType `json:"type"`
		Attributes struct {
			ReferenceID  string `json:"referenceId"`
			Instructions struct {
				DisplayName string `json:"displayName"`

				// VA.
				BankShortCode BankCode `json:"bankShortCode"`
				AccountNo     string   `json:"accountNo"`

				// QRIS.
				ImageURL string `json:"imageUrl"`
			} `json:"instructions"`
		} `json:"attributes"`
	} `json:"data"`
}

func (p paymentMethod) toPaymentMethod() *PaymentMethod {
	return &PaymentMethod{
		ID:            p.Data.ID,
		Type:          p.Data.Type,
		ReferenceID:   p.Data.Attributes.ReferenceID,
		DisplayName:   p.Data.Attributes.Instructions.DisplayName,
		BankShortCode: p.Data.Attributes.Instructions.BankShortCode,
		AccountNo:     p.Data.Attributes.Instructions.AccountNo,
		ImageURL:      p.Data.Attributes.Instructions.ImageURL,
	}
}

type paymentMethodAction struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			TargetID   string `json:"targetId"`
			TargetType string `json:"targetType"`
			Name       Action `json:"action"`
			Options    struct {
				Amount string `json:"amount"`
			} `json:"options"`
		} `json:"attributes"`
	} `json:"data"`
}

func (p paymentMethodAction) toPaymentMethodAction() *PaymentMethodAction {
	return &PaymentMethodAction{
		TargetID:   p.Data.Attributes.TargetID,
		TargetType: p.Data.Attributes.TargetType,
		Name:       p.Data.Attributes.Name,
	}
}
