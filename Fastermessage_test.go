package fastermessage

import (
	"log"
	"testing"
)

func TestSendFastermessageSms(t *testing.T) {

	fasterAccount := NewFasterAccount("")

	result, err := fasterAccount.SendSms("Faster", "22996034411", "Hello, Faster! ")
	if err != nil {
		t.Error("Error Send sms ", err)
	}
	log.Println("[INFO], TestSendFastermessageSms ,  ", result)
}
