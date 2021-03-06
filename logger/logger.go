package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Handler interface {
	Post(title, message, channel string) error
}

type handlerNS struct {
	url string
}

type Level string

const (
	Info      = Level("Info")
	Error     = Level("Error")
	Debug     = Level("Debug")
	Alert     = Level("Alert")
	Warning   = Level("Warning")
	Notice    = Level("Notice")
	Critical  = Level("Critical")
	Emergency = Level("Emergency")
)

type Logger interface {
	Log(title, message, channel string) error
	Critical(message string) error
	Debug(message string) error
	Error(message string) error
	Info(message string) error
	Warning(message string) error
}

type logger struct {
	handlers []Handler
}

func (l Level) IsValid() bool {
	for _, level := range []Level{Info, Error, Debug, Alert, Warning, Notice, Critical, Emergency} {
		if l == level {
			return true
		}
	}
	return false
}

func (l Level) Name() string {
	return strings.ToUpper(string(l))
}

func NewLogger(handlers []Handler) Logger {
	return logger{handlers: handlers}
}

func (al logger) Log(title, message, channel string) error {
	for _, handler := range al.handlers {
		err := handler.Post(title, message, channel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (al logger) Info(message string) error {
	return al.Log(string(Info), message, "")
}

func (al logger) Error(message string) error {
	return al.Log(string(Error), message, "")
}

func (al logger) Debug(message string) error {
	return al.Log(string(Debug), message, "")
}

func (al logger) Warning(message string) error {
	return al.Log(string(Warning), message, "")
}

func (al logger) Critical(message string) error {
	return al.Log(string(Critical), message, "")
}

func NewHandlerNS(url string) handlerNS {
	return handlerNS{url: url}
}

func (h handlerNS) Post(title, message, channel string) error {
	body, err := h.getBody(title, message, channel)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, h.url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpClient := &http.Client{Timeout: 5 * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Error log to %s. Request Status: %s", h.url, res.Status))
	}
	return nil
}

func (h handlerNS) getBody(title, message, channel string) ([]byte, error) {
	log := map[string]string{
		"origin":  channel,
		"title":   title,
		"message": message,
	}
	content := map[string]interface{}{
		"type": []string{"log"},
		"log":  log,
	}
	return json.Marshal(content)
}
