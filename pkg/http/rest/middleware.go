package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/valentyn88/password_service"
)

const (
	reqParamKey = "req_param"

	errInvalidParamsStr   = "invalid parameters were passed"
	errInvalidTotalLenStr = "total length of params can't be less than: %d and more than %d"
	errInvalidCountStr    = "param count can't be more than %d"
)

// ErrorMsg describes error message.
type ErrorMsg struct {
	Error string `json:"error"`
}

// RequestParam middleware for incoming parameters.
func (h Handler) RequestParam(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("couldn't read request body error: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Println(err.Error())
			}
		}()

		var reqParam password_service.RequestParam
		if err := json.Unmarshal(body, &reqParam); err != nil {
			log.Printf("invalid parameters error: %s\n", err)
			errMsg(w, errInvalidParamsStr, 0)
			return
		}

		reqParam.DefVal()

		if !password_service.IsValidReqParamsLen(reqParam) {
			log.Println("invalid total length of request params")
			errMsg(w, fmt.Sprintf(errInvalidTotalLenStr, password_service.MinTotalLen,
				password_service.MaxTotalLen), 0)
			return
		}

		if !password_service.IsValidReqParamsCount(reqParam) {
			log.Println("invalid count of request params")
			errMsg(w, fmt.Sprintf(errInvalidCountStr, password_service.MaxCount), 0)
			return
		}

		ctx := context.WithValue(r.Context(), reqParamKey, reqParam)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func errMsg(w http.ResponseWriter, s string, code int) {
	if code == 0 {
		code = http.StatusBadRequest
	}
	w.WriteHeader(code)
	b, err := json.Marshal(ErrorMsg{Error: s})
	if err != nil {
		log.Println(err.Error())
	}
	if _, err := w.Write(b); err != nil {
		log.Println(err.Error())
	}
}
