package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/valentyn88/password_service"
)

// PasswordResponse describes password reponse.
type PasswordResponse struct {
	Passwords []password_service.Password `json:"passwords"`
}

// Handler describes struct for handlers.
type Handler struct {
	PasswordService password_service.Creator
}

// GeneratePasswords creates passwords.
func (h Handler) GeneratePasswords(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqParam, ok := ctx.Value(reqParamKey).(password_service.RequestParam)
	if !ok {
		log.Println("couldn't get request params from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var pp []password_service.Password
	for i := 0; i < reqParam.Count; i++ {
		pp = append(pp, h.PasswordService.Create(reqParam.LettersLen, reqParam.SpecChLen, reqParam.NumbersLen))
	}
	resp, err := json.Marshal(PasswordResponse{Passwords: pp})
	if err != nil {
		log.Printf("couldn't marshal products response error: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("couldn't write response error: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
