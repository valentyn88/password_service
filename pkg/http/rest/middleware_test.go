package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"strings"

	"github.com/valentyn88/password_service/mock"
)

func TestHandler_RequestParam(t *testing.T) {
	testCases := []struct {
		name           string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "invalid params",
			body:           "test",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid parameters were passed"}`,
		},
		{
			name:           "invalid total length of password",
			body:           `{"letters_len":10,"spec_ch_len":2,"numbers_len":4, "count":10}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"total length of params can't be less than: 5 and more than 15"}`,
		},
		{
			name:           "invalid count of passwords",
			body:           `{"letters_len":10,"spec_ch_len":1,"numbers_len":4, "count":51}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"param count can't be more than 50"}`,
		},
	}

	var h Handler
	h.PasswordService = mock.PasswordSvc{}

	handler := http.HandlerFunc(h.RequestParam(h.GeneratePasswords))

	for _, tc := range testCases {
		t.Log(tc.name)

		bb := bytes.NewBuffer([]byte(tc.body))
		req, err := http.NewRequest("POST", "/v1/generate", bb)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != tc.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusBadRequest)
		}

		if tc.expectedBody != "" {
			got := strings.TrimRight(rr.Body.String(), "\r\n")
			if got != tc.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					got, tc.expectedBody)
			}
		}
	}
}
