package fastermessage

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//FasterAccount : Account Object
type FasterAccount struct {
	XapiKey string
}

//ResponseFasterSms :
type ResponseFasterSms struct {
	Status        bool   `json:"status"`
	From          string `json:"from"`
	To            string `json:"to"`
	Text          string `json:"text"`
	Code          int    `json:"code"`
	Messagestatus string `json:"messagestatus"`
	Description   string `json:"description"`

	SmsCount  int    `json:"smsCount"`
	Devise    string `json:"devise"`
	UnitPrice string `json:"unitPrice"`

	MessagePrice int    `json:"messagePrice"`
	MessageID    string `json:"messageId"`
}

//NewFasterAccount : To create a new Faster msg acount
func NewFasterAccount(XapiKey string) FasterAccount {
	return FasterAccount{XapiKey: XapiKey}
}

//SendSms :Send Sms
func (f *FasterAccount) SendSms(from, to, msg string) (ResponseFasterSms, error) {
	var respObj ResponseFasterSms

	endpoint := "https://api.fastermessage.com/v1/sms/send"
	data := url.Values{}
	data.Set("from", from)
	data.Set("to", to)
	data.Set("text", msg)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.Header.Add("X-API-KEY", f.XapiKey)

	res, err := client.Do(r)
	if err != nil {
		return respObj, err
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&respObj)
	if err != nil {
		return respObj, err
	}
	return respObj, nil
}
