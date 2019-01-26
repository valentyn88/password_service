package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/valentyn88/password_service/pkg/http/rest"
	"github.com/valentyn88/password_service/pkg/secure_password"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	log.SetOutput(os.Stdout)

	var h rest.Handler
	h.PasswordService = secure_password.PasswordService{}

	http.HandleFunc("/v1/generate", h.RequestParam(h.GeneratePasswords))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("couldn't start password server error: %s", err.Error())
	}
}
