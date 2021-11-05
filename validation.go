package xfers

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

var mod *mold.Transformer
var val *validator.Validate

func init() {
	val = validator.New()
	val.RegisterValidationCtx("bank_code", validateBankCode)
	val.RegisterValidationCtx("status", validateStatus)
	val.RegisterValidationCtx("payment_action", validationPaymentAction)
	val.RegisterValidationCtx("disbursement_action", validationDisbursementAction)
	val.RegisterValidationCtx("payment_type", validationPaymentType)
	val.RegisterValidationCtx("retail_outlet", validationRetailOutlet)
	val.RegisterValidationCtx("va_bank_code", validationVABankCode)
	val.RegisterValidationCtx("e_wallet", validationEWallet)

	mod = modifiers.New()
	mod.Register("no_space", modNoSpace)
}

func modNoSpace(ctx context.Context, fl mold.FieldLevel) error {
	switch fl.Field().Kind() {
	case reflect.String:
		fl.Field().SetString(strings.Replace(fl.Field().String(), " ", "", -1))
	}
	return nil
}

func validate(data interface{}) error {
	if err := mod.Struct(context.Background(), data); err != nil {
		return err
	}
	if err := val.Struct(data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				return errRequiredField(e.Field())
			case "gt":
				return errGTField(e.Field(), e.Param())
			case "gte":
				return errGTEField(e.Field(), e.Param())
			case "max":
				return errMaxField(e.Field(), e.Param())
			case "numeric":
				return errNumericField(e.Field())
			case "url":
				return errURLField(e.Field())
			default:
				return errInvalidValueField(e.Field())
			}
		}
		return err
	}
	return nil
}

func validateStatus(ctx context.Context, fl validator.FieldLevel) bool {
	return map[Status]bool{
		"":               true,
		StatusPending:    true,
		StatusProcessing: true,
		StatusCancelled:  true,
		StatusExpired:    true,
		StatusFailed:     true,
		StatusPaid:       true,
		StatusCompleted:  true,
	}[Status(fl.Field().String())]
}

func validationPaymentAction(ctx context.Context, fl validator.FieldLevel) bool {
	return map[Action]bool{
		ActionCancel:         true,
		ActionReceivePayment: true,
		ActionSettle:         true,
	}[Action(fl.Field().String())]
}

func validationDisbursementAction(ctx context.Context, fl validator.FieldLevel) bool {
	return map[Action]bool{
		ActionComplete: true,
		ActionFail:     true,
	}[Action(fl.Field().String())]
}

func validationPaymentType(ctx context.Context, fl validator.FieldLevel) bool {
	return map[PaymentType]bool{
		PaymentVA:      true,
		PaymentOutlet:  true,
		PaymentEWallet: true,
		PaymentQRIS:    true,
	}[PaymentType(fl.Field().String())]
}

func validationRetailOutlet(ctx context.Context, fl validator.FieldLevel) bool {
	return map[RetailOutlet]bool{
		OutletAlfamart:  true,
		OutletIndomaret: true,
	}[RetailOutlet(fl.Field().String())]
}

func validationVABankCode(ctx context.Context, fl validator.FieldLevel) bool {
	return map[BankCode]bool{
		BankBCA:              true,
		BankBRI:              true,
		BankBNI:              true,
		BankMandiri:          true,
		BankCIMB:             true,
		BankDanamon:          true,
		BankPermata:          true,
		BankHana:             true,
		BankSahabatSampoerna: true,
	}[BankCode(fl.Field().String())]
}

func validationEWallet(ctx context.Context, fl validator.FieldLevel) bool {
	return map[EWallet]bool{
		EWalletShopeePay: true,
	}[EWallet(fl.Field().String())]
}

func validateBankCode(ctx context.Context, fl validator.FieldLevel) bool {
	return map[BankCode]bool{
		BankBCA:                          true,
		BankMandiri:                      true,
		BankBNI:                          true,
		BankPermata:                      true,
		BankBRI:                          true,
		BankCIMB:                         true,
		BankDanamon:                      true,
		BankPanin:                        true,
		BankMaybank:                      true,
		BankAnglomas:                     true,
		BankBangkok:                      true,
		BankAgris:                        true,
		BankSinarmas:                     true,
		BankAgroniaga:                    true,
		BankAndara:                       true,
		BankAntarDaerah:                  true,
		BankANZ:                          true,
		BankArtha:                        true,
		BankArtos:                        true,
		BankBisnis:                       true,
		BankBJB:                          true,
		BankBNP:                          true,
		BankBukopin:                      true,
		BankBumiArta:                     true,
		BankCapital:                      true,
		BankBCASyariah:                   true,
		BankChinatrus:                    true,
		BankCIMBUSS:                      true,
		BankCommonwealth:                 true,
		BankDanamonUUS:                   true,
		BankDBS:                          true,
		BankDinar:                        true,
		BankDKI:                          true,
		BankDKIUSS:                       true,
		BankEkonomi:                      true,
		BankFama:                         true,
		BankGanesha:                      true,
		BankHana:                         true,
		BankHarda:                        true,
		BankHimpunanSaudara:              true,
		BankICBC:                         true,
		BankInaPerdana:                   true,
		BankIndexSelindo:                 true,
		BankJasaJakarta:                  true,
		BankKesejahteraanEkonomi:         true,
		BankMaspion:                      true,
		BankMayapada:                     true,
		BankMaybankSyariah:               true,
		BankMayora:                       true,
		BankMega:                         true,
		BankMestikaDharma:                true,
		BankMetroExpress:                 true,
		BankMizuho:                       true,
		BankMNC:                          true,
		BankMuamalat:                     true,
		BankMultiArtaSentosa:             true,
		BankMutiara:                      true,
		BankNationalnobu:                 true,
		BankNusantaraParahyangan:         true,
		BankOCBC:                         true,
		BankOCBCUUS:                      true,
		BankBAML:                         true,
		BankBOC:                          true,
		BankIndia:                        true,
		BankTokyo:                        true,
		BankPaninSyariah:                 true,
		BankPermataUUS:                   true,
		BankPundi:                        true,
		BankQNBKesawan:                   true,
		BankRabobank:                     true,
		BankResona:                       true,
		BankRoyal:                        true,
		BankSahabatPurbaDanarta:          true,
		BankSahabatSampoerna:             true,
		BankSBI:                          true,
		BankSinarHarapanBali:             true,
		BankMitsui:                       true,
		BankBRISyariah:                   true,
		BankBukopinSyariah:               true,
		BankMandiriSyariah:               true,
		BankMegaSyariah:                  true,
		BankBTN:                          true,
		BankBTNUUS:                       true,
		BankTabunganPensiunanNasional:    true,
		BankTabunganPensiunanNasionalUUS: true,
		BankUOB:                          true,
		BankVictoria:                     true,
		BankVictoriaSyariah:              true,
		BankWindu:                        true,
		BankWoori:                        true,
		BankYudhaBhakti:                  true,
		BankAceh:                         true,
		BankAcehUUS:                      true,
		BankBali:                         true,
		BankBengkulu:                     true,
		BankBPDDIY:                       true,
		BankBPDDIYSyariah:                true,
		BankJambi:                        true,
		BankJambiUUS:                     true,
		BankJawaTengah:                   true,
		BankJawaTengahUUS:                true,
		BankJawaTimur:                    true,
		BankJawaTimurUUS:                 true,
		BankKalimantanBarat:              true,
		BankKalimantanBaratUUS:           true,
		BankKalimantanSelatan:            true,
		BankKalimantanSelatanUUS:         true,
		BankKalimatanTengah:              true,
		BankKalimatanTimur:               true,
		BankKalimantanTimurUUS:           true,
		BankLampung:                      true,
		BankMaluku:                       true,
		BankNusaTenggaraBarat:            true,
		BankNusaTenggaraBaratUUS:         true,
		BankNusaTenggaraTimur:            true,
		BankPapua:                        true,
		BankRiauKepri:                    true,
		BankRiaouKepriUUS:                true,
		BankSulawesi:                     true,
		BankSulawesiTenggara:             true,
		BankSulselbar:                    true,
		BankSulselbarUUS:                 true,
		BankSulut:                        true,
		BankSumateraBarat:                true,
		BankSumateraBaratUUS:             true,
		BankSumselBabel:                  true,
		BankSumselBabelUUS:               true,
		BankSumut:                        true,
		BankSumutUUS:                     true,
		BankCentratama:                   true,
		BankCitibank:                     true,
		BankDeutsche:                     true,
		BankHSBC:                         true,
		BankHSBCUUS:                      true,
		BankJPMorgan:                     true,
		BankPrimaMaster:                  true,
		BankStandardCharted:              true,
		BankMitraNiaga:                   true,
		BankEkspor:                       true,
		BankArtaNiagaKencana:             true,
		BankBJBSyariah:                   true,
		BankBNISyariah:                   true,
	}[BankCode(fl.Field().String())]
}
