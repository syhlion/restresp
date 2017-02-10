package restresp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type customError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (c *customError) Error() string {
	return c.Msg
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	d := &customError{5, "wtf"}
	Write(w, d, 200)
}

func TestWrite(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":{"code":5,"msg":"wtf"}}`
	body := strings.TrimSpace(rr.Body.String())
	if body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expected)
	}
}
