package xfers

// EnvironmentType is type for environment.
type EnvironmentType int8

// Available options for EnvironmentType.
const (
	Sandbox EnvironmentType = iota
	Production
)

var envURL = map[EnvironmentType]string{
	Sandbox:    "https://sandbox-id.xfers.com/api/v4",
	Production: "https://id.xfers.com/api/v4",
}

var envLog = map[EnvironmentType]LogLevel{
	Sandbox:    LogDebug,
	Production: LogError,
}

// Status is type for payment & disbursement transaction status.
type Status string

// Available options for Status.
const (
	StatusPending    Status = "pending"    // payment & disbursement
	StatusProcessing Status = "processing" // payment & disbursement
	StatusCancelled  Status = "cancelled"  // payment
	StatusExpired    Status = "expired"    // payment
	StatusFailed     Status = "failed"     // payment & disbursement
	StatusPaid       Status = "paid"       // payment
	StatusCompleted  Status = "completed"  // payment & disbursement
)

// Action is type for payment & disbursement simulate action.
type Action string

// Available options for Action.
const (
	ActionCancel         Action = "cancel"          // payment
	ActionReceivePayment Action = "receive_payment" // payment
	ActionSettle         Action = "settle"          // payment
	ActionComplete       Action = "complete"        // disbursement
	ActionFail           Action = "fail"            // disbursement
)

// PaymentType is type for payment method.
type PaymentType string

// Available options for PaymentType.
const (
	PaymentVA      PaymentType = "virtual_bank_account"
	PaymentOutlet  PaymentType = "retail_outlet"
	PaymentQRIS    PaymentType = "qris"
	PaymentEWallet PaymentType = "e-wallet"
)

// RetailOutlet is type for retail outlet.
type RetailOutlet string

// Available options for RetailOutlet.
const (
	OutletAlfamart  RetailOutlet = "ALFAMART"
	OutletIndomaret RetailOutlet = "INDOMARET"
)

// EWallet is type for e-wallet.
type EWallet string

// Available options for e-wallet.
const (
	EWalletShopeePay EWallet = "SHOPEEPAY"
)

// BankCode is code for banks.
type BankCode string

// Available options for BankCode.
const (
	BankBCA                          BankCode = "BCA"        // va persist & one
	BankMandiri                      BankCode = "MANDIRI"    // va persist & one
	BankBNI                          BankCode = "BNI"        // va persist & one
	BankPermata                      BankCode = "PERMATA"    // va persist
	BankBRI                          BankCode = "BRI"        // va persist & one
	BankCIMB                         BankCode = "CIMB_NIAGA" // va persist
	BankDanamon                      BankCode = "DANAMON"    // va persist
	BankPanin                        BankCode = "PANIN"
	BankMaybank                      BankCode = "BII"
	BankAnglomas                     BankCode = "ANGLOMAS"
	BankBangkok                      BankCode = "BANGKOK"
	BankAgris                        BankCode = "AGRIS"
	BankSinarmas                     BankCode = "SINARMAS"
	BankAgroniaga                    BankCode = "AGRONIAGA"
	BankAndara                       BankCode = "ANDARA"
	BankAntarDaerah                  BankCode = "ANTAR_DAERAH"
	BankANZ                          BankCode = "ANZ"
	BankArtha                        BankCode = "ARTHA"
	BankArtos                        BankCode = "ARTOS"
	BankBisnis                       BankCode = "BISNIS"
	BankBJB                          BankCode = "BJB"
	BankBNP                          BankCode = "BNP"
	BankBukopin                      BankCode = "BUKOPIN"
	BankBumiArta                     BankCode = "BUMI_ARTA"
	BankCapital                      BankCode = "CAPITAL"
	BankBCASyariah                   BankCode = "BCA_SYR"
	BankChinatrus                    BankCode = "CHINATRUST"
	BankCIMBUSS                      BankCode = "CIMB_UUS"
	BankCommonwealth                 BankCode = "COMMONWEALTH"
	BankDanamonUUS                   BankCode = "DANAMON_UUS"
	BankDBS                          BankCode = "DBS"
	BankDinar                        BankCode = "DINAR_INDONESIA"
	BankDKI                          BankCode = "DKI"
	BankDKIUSS                       BankCode = "DKI_UUS"
	BankEkonomi                      BankCode = "EKONOMI_RAHARJA"
	BankFama                         BankCode = "FAMA"
	BankGanesha                      BankCode = "GANESHA"
	BankHana                         BankCode = "HANA" // va persist
	BankHarda                        BankCode = "HARDA_INTERNASIONAL"
	BankHimpunanSaudara              BankCode = "HIMPUNAN_SAUDARA"
	BankICBC                         BankCode = "ICBC"
	BankInaPerdana                   BankCode = "INA_PERDANA"
	BankIndexSelindo                 BankCode = "INDEX_SELINDO"
	BankJasaJakarta                  BankCode = "JASA_JAKARTA"
	BankKesejahteraanEkonomi         BankCode = "KESEJAHTERAAN_EKONOMI"
	BankMaspion                      BankCode = "MASPION"
	BankMayapada                     BankCode = "MAYAPADA"
	BankMaybankSyariah               BankCode = "MAYBANK_SYR"
	BankMayora                       BankCode = "MAYORA"
	BankMega                         BankCode = "MEGA"
	BankMestikaDharma                BankCode = "MESTIKA_DHARMA"
	BankMetroExpress                 BankCode = "METRO_EXPRESS"
	BankMizuho                       BankCode = "MIZUHO"
	BankMNC                          BankCode = "MNC_INTERNASIONAL"
	BankMuamalat                     BankCode = "MUAMALAT"
	BankMultiArtaSentosa             BankCode = "MULTI_ARTA_SENTOSA"
	BankMutiara                      BankCode = "MUTIARA"
	BankNationalnobu                 BankCode = "NATIONALNOBU"
	BankNusantaraParahyangan         BankCode = "NUSANTARA_PARAHYANGAN"
	BankOCBC                         BankCode = "OCBC"
	BankOCBCUUS                      BankCode = "OCBC_UUS"
	BankBAML                         BankCode = "BAML"
	BankBOC                          BankCode = "BOC"
	BankIndia                        BankCode = "INDIA"
	BankTokyo                        BankCode = "TOKYO"
	BankPaninSyariah                 BankCode = "PANIN_SYR"
	BankPermataUUS                   BankCode = "PERMATA_UUS"
	BankPundi                        BankCode = "PUNDI_INDONESIA"
	BankQNBKesawan                   BankCode = "QNB_KESAWAN"
	BankRabobank                     BankCode = "RABOBANK"
	BankResona                       BankCode = "RESONA"
	BankRoyal                        BankCode = "ROYAL"
	BankSahabatPurbaDanarta          BankCode = "SAHABAT_PURBA_DANARTA"
	BankSahabatSampoerna             BankCode = "SAHABAT_SAMPOERNA" // va persist & one
	BankSBI                          BankCode = "SBI_INDONESIA"
	BankSinarHarapanBali             BankCode = "SINAR_HARAPAN_BALI"
	BankMitsui                       BankCode = "MITSUI"
	BankBRISyariah                   BankCode = "BRI_SYR"
	BankBukopinSyariah               BankCode = "BUKOPIN_SYR"
	BankMandiriSyariah               BankCode = "MANDIRI_SYR"
	BankMegaSyariah                  BankCode = "MEGA_SYR"
	BankBTN                          BankCode = "BTN"
	BankBTNUUS                       BankCode = "BTN_UUS"
	BankTabunganPensiunanNasional    BankCode = "TABUNGAN_PENSIUNAN_NASIONAL"
	BankTabunganPensiunanNasionalUUS BankCode = "TABUNGAN_PENSIUNAN_NASIONAL_UUS"
	BankUOB                          BankCode = "UOB"
	BankVictoria                     BankCode = "VICTORIA_INTERNASIONAL"
	BankVictoriaSyariah              BankCode = "VICTORIA_SYR"
	BankWindu                        BankCode = "WINDU"
	BankWoori                        BankCode = "WOORI"
	BankYudhaBhakti                  BankCode = "YUDHA_BHAKTI"
	BankAceh                         BankCode = "ACEH"
	BankAcehUUS                      BankCode = "ACEH_UUS"
	BankBali                         BankCode = "BALI"
	BankBengkulu                     BankCode = "BENGKULU"
	BankBPDDIY                       BankCode = "BPD_DIY"
	BankBPDDIYSyariah                BankCode = "BPD_DIY_SYR"
	BankJambi                        BankCode = "JAMBI"
	BankJambiUUS                     BankCode = "JAMBI_UUS"
	BankJawaTengah                   BankCode = "JAWA_TENGAH"
	BankJawaTengahUUS                BankCode = "JAWA_TENGAH_UUS"
	BankJawaTimur                    BankCode = "JAWA_TIMUR"
	BankJawaTimurUUS                 BankCode = "JAWA_TIMUR_UUS"
	BankKalimantanBarat              BankCode = "KALIMANTAN_BARAT"
	BankKalimantanBaratUUS           BankCode = "KALIMANTAN_BARAT_UUS"
	BankKalimantanSelatan            BankCode = "KALIMANTAN_SELATAN"
	BankKalimantanSelatanUUS         BankCode = "KALIMANTAN_SELATAN_UUS"
	BankKalimatanTengah              BankCode = "KALIMANTAN_TENGAH"
	BankKalimatanTimur               BankCode = "KALIMANTAN_TIMUR"
	BankKalimantanTimurUUS           BankCode = "KALIMANTAN_TIMUR_UUS"
	BankLampung                      BankCode = "LAMPUNG"
	BankMaluku                       BankCode = "MALUKU"
	BankNusaTenggaraBarat            BankCode = "NUSA_TENGGARA_BARAT"
	BankNusaTenggaraBaratUUS         BankCode = "NUSA_TENGGARA_BARAT_UUS"
	BankNusaTenggaraTimur            BankCode = "NUSA_TENGGARA_TIMUR"
	BankPapua                        BankCode = "PAPUA"
	BankRiauKepri                    BankCode = "RIAU_DAN_KEPRI"
	BankRiaouKepriUUS                BankCode = "RIAU_DAN_KEPRI_UUS"
	BankSulawesi                     BankCode = "SULAWESI"
	BankSulawesiTenggara             BankCode = "SULAWESI_TENGGARA"
	BankSulselbar                    BankCode = "SULSELBAR"
	BankSulselbarUUS                 BankCode = "SULSELBAR_UUS"
	BankSulut                        BankCode = "SULUT"
	BankSumateraBarat                BankCode = "SUMATERA_BARAT"
	BankSumateraBaratUUS             BankCode = "SUMATERA_BARAT_UUS"
	BankSumselBabel                  BankCode = "SUMSEL_DAN_BABEL"
	BankSumselBabelUUS               BankCode = "SUMSEL_DAN_BABEL_UUS"
	BankSumut                        BankCode = "SUMUT"
	BankSumutUUS                     BankCode = "SUMUT_UUS"
	BankCentratama                   BankCode = "CENTRATAMA"
	BankCitibank                     BankCode = "CITIBANK"
	BankDeutsche                     BankCode = "DEUTSCHE"
	BankHSBC                         BankCode = "HSBC"
	BankHSBCUUS                      BankCode = "HSBC_UUS"
	BankJPMorgan                     BankCode = "JPMORGAN"
	BankPrimaMaster                  BankCode = "PRIMA_MASTER"
	BankStandardCharted              BankCode = "STANDARD_CHARTERED"
	BankMitraNiaga                   BankCode = "MITRA_NIAGA"
	BankEkspor                       BankCode = "EKSPOR_INDONESIA"
	BankArtaNiagaKencana             BankCode = "ARTA_NIAGA_KENCANA"
	BankBJBSyariah                   BankCode = "BJB_SYR"
	BankBNISyariah                   BankCode = "BNI_SYR"
)
