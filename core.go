package nicepay

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
)

type CoreGateway struct {
	Client Client
}

func (gateway *CoreGateway) Registration(req *RegistrationRequest) (Response, error) {
	req.IMID = gateway.Client.MID
	req.DBProcessURL = gateway.Client.NotifURL

	path := gateway.Client.APIEnvType.String() + "/nicepay/direct/v2/registration"
	resp := Response{}
	jsonReq, _ := json.Marshal(req)

	headers := []Header{
		Header{Key: "Content-Type", Value: "application/json"},
	}

	err := gateway.Client.CallRequest("POST", path, headers, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (gateway *CoreGateway) Payment(w http.ResponseWriter, req *PaymentRequest) error {
	req.CallBackURL = gateway.Client.CallbackURL

	path := gateway.Client.APIEnvType.String() + "/nicepay/direct/v2/payment"

	t, err := template.ParseFiles("go-nicepay/payment.html")
	if err != nil {
		return err
	}

	inputs := []Input{
		Input{Name: "timeStamp", Value: req.TimeStamp},
		Input{Name: "tXid", Value: req.TXID},
		Input{Name: "merchantToken", Value: req.MerchantToken},
		Input{Name: "cardNo", Value: req.CardNo},
		Input{Name: "cardExpYymm", Value: req.CardExpYymm},
		Input{Name: "cardCvv", Value: req.CardCvv},
		Input{Name: "cardHolderNm", Value: req.CardHolderNm},
		Input{Name: "recurringToken", Value: req.RecurringToken},
		Input{Name: "preauthToken", Value: req.PreauthToken},
		Input{Name: "clickPayNo", Value: req.ClickPayNo},
		Input{Name: "dataField3", Value: req.DataField3},
		Input{Name: "clickPayToken", Value: req.ClickPayToken},
		Input{Name: "callBackUrl", Value: req.CallBackURL},
	}

	request := PostPaymentRequest{
		URL:    path,
		Inputs: inputs,
	}

	err = t.Execute(w, request)
	if err != nil {
		return err
	}

	return nil
}

func (gateway *CoreGateway) GetOrderStatus(req *OrderStatusRequest) (Response, error) {
	req.IMID = gateway.Client.MID

	path := gateway.Client.APIEnvType.String() + "/nicepay/direct/v2/inquiry"
	resp := Response{}
	jsonReq, _ := json.Marshal(req)

	headers := []Header{
		Header{Key: "Content-Type", Value: "application/json"},
	}

	err := gateway.Client.CallRequest("POST", path, headers, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (gateway *CoreGateway) GenerateMerchantToken(req *Token) string {
	req.IMID = gateway.Client.MID
	req.MerchantKey = gateway.Client.MKey

	str := req.TimeStamp + req.IMID + req.ReferenceNo + req.Amt + req.MerchantKey
	token := Sha256(str)

	return token
}
