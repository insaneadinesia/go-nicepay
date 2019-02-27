package nicepay

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Input struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PostPaymentRequest struct {
	URL    string  `json:"url"`
	Inputs []Input `json:"inputs"`
}

type RegistrationRequest struct {
	TimeStamp       string `json:"timeStamp"`
	IMID            string `json:"iMid"`
	PayMethod       string `json:"payMethod"`
	Currency        string `json:"currency"`
	Amt             string `json:"amt"`
	ReferenceNo     string `json:"referenceNo"`
	GoodsNm         string `json:"goodsNm"`
	BillingNm       string `json:"billingNm"`
	BillingPhone    string `json:"billingPhone"`
	BillingEmail    string `json:"billingEmail"`
	BillingAddr     string `json:"billingAddr"`
	BillingCity     string `json:"billingCity"`
	BillingState    string `json:"billingState"`
	BillingPostCd   string `json:"billingPostCd"`
	BillingCountry  string `json:"billingCountry"`
	DeliveryNm      string `json:"deliveryNm"`
	DeliveryPhone   string `json:"deliveryPhone"`
	DeliveryAddr    string `json:"deliveryAddr"`
	DeliveryCity    string `json:"deliveryCity"`
	DeliveryState   string `json:"deliveryState"`
	DeliveryPostCd  string `json:"deliveryPostCd"`
	DeliveryCountry string `json:"deliveryCountry"`
	DBProcessURL    string `json:"dbProcessUrl"`
	Vat             string `json:"vat"`
	Fee             string `json:"fee"`
	NotaxAmt        string `json:"notaxAmt"`
	Description     string `json:"description"`
	MerchantToken   string `json:"merchantToken"`
	ReqDt           string `json:"reqDt"`
	ReqTm           string `json:"reqTm"`
	ReqDomain       string `json:"reqDomain"`
	ReqServerIP     string `json:"reqServerIP"`
	ReqClientVer    string `json:"reqClientVer"`
	UserIP          string `json:"userIP"`
	UserSessionID   string `json:"userSessionID"`
	UserAgent       string `json:"userAgent"`
	UserLanguage    string `json:"userLanguage"`
	CartData        string `json:"cartData"`
	InstmntType     string `json:"instmntType"`
	InstmntMon      string `json:"instmntMon"`
	RecurrOpt       string `json:"recurrOpt"`
	BankCd          string `json:"bankCd"`
	VacctValidDt    string `json:"vacctValidDt"`
	VacctValidTm    string `json:"vacctValidTm"`
	MerFixAcctID    string `json:"merFixAcctId"`
	MitraCd         string `json:"mitraCd"`
	MRefNo          string `json:"mRefNo"`
	PayValidDt      string `json:"payValidDt"`
	PayValidTm      string `json:"payValidTm"`
}

type PaymentRequest struct {
	TimeStamp      string `json:"timeStamp"`
	TXID           string `json:"tXid"`
	MerchantToken  string `json:"merchantToken"`
	CardNo         string `json:"cardNo"`
	CardExpYymm    string `json:"cardExpYymm"`
	CardCvv        string `json:"cardCvv"`
	CardHolderNm   string `json:"cardHolderNm"`
	RecurringToken string `json:"recurringToken"`
	PreauthToken   string `json:"preauthToken"`
	ClickPayNo     string `json:"clickPayNo"`
	DataField3     string `json:"dataField3"`
	ClickPayToken  string `json:"clickPayToken"`
	CallBackURL    string `json:"callBackUrl"`
}

type NotificationRequest struct {
	TXID          string `json:"tXid"`
	ReferenceNo   string `json:"referenceNo"`
	Amt           string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
	MatchCl       string `json:"matchCl"`
	Status        string `json:"status"`
	BankCd        string `json:"bankCd"`
	VacctNo       string `json:"vacctNo"`
	AuthNo        string `json:"authNo"`
	CardNo        string `json:"cardNo"`
	IssuBankCd    string `json:"issuBankCd"`
	IssuBankNm    string `json:"issuBankNm"`
	AcquBankCd    string `json:"acquBankCd"`
	AcquBankNm    string `json:"acquBankNm"`
	DepositDt     string `json:"depositDt"`
	DepositTm     string `json:"depositTm"`
	PayNo         string `json:"payNo"`
	MitraCd       string `json:"mitraCd"`
}

type OrderStatusRequest struct {
	TimeStamp     string `json:"timeStamp"`
	TXID          string `json:"tXid"`
	IMID          string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo"`
	Amt           string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}

type Token struct {
	TimeStamp   string `json:"timeStamp"`
	IMID        string `json:"iMid"`
	ReferenceNo string `json:"referenceNo"`
	Amt         string `json:"amt"`
	MerchantKey string `json:"merchantKey"`
}
