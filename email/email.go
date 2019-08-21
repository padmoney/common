package email

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/padmoney/common/ns"
)

type Sender interface {
	Send(from string, to []string, subject string, content string) error
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

func (h SenderNS) Send(from string, to []string, subject string, content string) error {
	if h.dev {
		t := "Envio de E-mail: \n\tFrom: %s\n\tTo: %s\n\tSubject: %s\n\tContent: %s"
		fmt.Println(fmt.Sprintf(t, from, strings.Join(to, ", "), subject, content))
		return nil
	}

	data := map[string]interface{}{
		"type": []string{"email"},
		"email": map[string]interface{}{
			"from":    from,
			"to":      to,
			"subject": subject,
			"content": content,
		},
	}
	encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return h.ns.SendNotification(encodedData)
}
