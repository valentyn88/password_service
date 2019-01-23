package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"context"

	"strings"

	"github.com/valentyn88/password_service"
	"github.com/valentyn88/password_service/mock"
)

func TestHandler_GeneratePasswords_Fail(t *testing.T) {
	req, err := http.NewRequest("POST", "/generate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	var h Handler
	h.PasswordService = mock.PasswordSvc{}

	handler := http.HandlerFunc(h.GeneratePasswords)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
	}
}

func TestHandler_GeneratePasswords_Success(t *testing.T) {
	req, err := http.NewRequest("POST", "/generate", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockReqParams := password_service.RequestParam{LettersLen: 2, SpecChLen: 2, NumbersLen: 5, Count: 1}

	ctx := req.Context()
	ctx = context.WithValue(ctx, reqParamKey, mockReqParams)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	var h Handler
	h.PasswordService = mock.PasswordSvc{}

	handler := http.HandlerFunc(h.GeneratePasswords)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	expected := `{"passwords":["8G628:z0\u0026"]}`
	got := strings.TrimRight(string(rr.Body.String()), "\r\n")
	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}
