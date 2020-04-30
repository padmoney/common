package credentials

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestNewCredentials(t *testing.T) {
	const (
		url  = "http://localhost"
		user = "mony"
		pass = "1234"
	)
	c := New(url, user, pass)
	token := c.Token()
	got, _ := base64.StdEncoding.DecodeString(token)
	expected := fmt.Sprintf("%s:%s", user, pass)
	if string(got) != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
	if c.URL() != url {
		t.Errorf("Expected %s, got %s", url, c.URL())
	}
}
