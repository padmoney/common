package ns

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/padmoney/common/credentials"
	"github.com/padmoney/common/rest"
)

func TestService_SendNotification(t *testing.T) {
	var body, method string
	fn := func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		method = r.Method
		body = string(b)
	}
	ts := httptest.NewServer(http.HandlerFunc(fn))
	cr := credentials.New(ts.URL, "padmoney", "123")
	cl := rest.NewClient(cr)
	ns := NewService(cl)
	data := map[string]interface{}{
		"types": []string{"email", "sms"},
		"email": map[string]interface{}{
			"from":    "from@test.com",
			"to":      []string{"to@test.com"},
			"subject": "test",
			"content": "test content",
		},
		"sms": map[string]interface{}{
			"phone":   "+5527999536698",
			"message": "content",
		},
	}
	encodedData, _ := json.Marshal(data)
	err := ns.SendNotification(encodedData)
	if err != nil {
		t.Error(err.Error())
	}
	if method != "POST" {
		t.Errorf("Expected '%s' got '%s'", "POST", method)
	}
	if body != string(encodedData) {
		t.Errorf("Expected '%s' got '%s'", encodedData, body)
	}
}

func TestService_SendNotificationGotError(t *testing.T) {
	var body, method string
	fn := func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		method = r.Method
		body = string(b)
		w.WriteHeader(400)
		w.Write([]byte("Error Message"))
	}
	ts := httptest.NewServer(http.HandlerFunc(fn))
	cr := credentials.New(ts.URL, "padmoney", "1234")
	cl := rest.NewClient(cr)
	ns := NewService(cl)
	data := map[string]interface{}{
		"types": []string{"email", "sms"},
		"email": map[string]interface{}{
			"from":    "from@test.com",
			"to":      []string{"to@test.com"},
			"subject": "test",
			"content": "test content",
		},
		"sms": map[string]interface{}{
			"phone":   "+5527999536698",
			"message": "content",
		},
	}
	encodedData, _ := json.Marshal(data)
	expectedErr := "Error Message"
	err := ns.SendNotification(encodedData)
	if err == nil {
		t.Error("Failed to assert error")
	}
	if method != "POST" {
		t.Errorf("Expected '%s' got '%s'", "POST", method)
	}
	if body != string(encodedData) {
		t.Errorf("Expected '%s' got '%s'", encodedData, body)
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected '%s' got '%s'", expectedErr, err.Error())
	}
}
