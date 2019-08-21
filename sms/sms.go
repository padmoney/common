package sms

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/padmoney/common/ns"
)

type Sender interface {
	Send(phone string, message string) error
}

type SenderNS struct {
	dev bool
	ns  ns.Service
}

func NewSender(dev bool, ns ns.Service) Sender {
	return SenderNS{
		dev: dev,
		ns:  ns,
	}
}

func (h SenderNS) Send(phone string, message string) error {
	if h.dev {
		t := "Envio de SMS: \n\tPhone: %s\n\tMessage: %s"
		fmt.Println(fmt.Sprintf(t, phone, message))
		return nil
	}

	data := map[string]interface{}{
		"type": []string{"sms"},
		"sms": map[string]interface{}{
			"phone":   formatPhone(phone),
			"message": message,
		},
	}
	encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return h.ns.SendNotification(encodedData)
}

func formatPhone(phone string) string {
	if strings.Contains(phone, "+") {
		return phone
	}
	return fmt.Sprintf("+55%s", phone)
}
