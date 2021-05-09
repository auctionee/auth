package handlers

import (
	"encoding/json"
	"github.com/auctionee/auth/pkg/data"
	"github.com/auctionee/auth/pkg/logger"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	l := logger.GetLogger(r.Context())
	creds := &data.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil{
		l.Fatal("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	l.Info(hashedPass)
}