package ns

import "github.com/padmoney/common/rest"

type Service struct {
	client rest.Client
}

func NewService(cl rest.Client) Service {
	return Service{cl}
}

func (s Service) SendNotification(data []byte) error {
	return s.client.Post("/notifications", data).Error
}
