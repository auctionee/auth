package handlers

import (
	"encoding/json"
	"github.com/auctionee/auth/pkg/data"
	"github.com/auctionee/auth/pkg/db"
	"github.com/auctionee/auth/pkg/logger"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger(r.Context())
	creds := &data.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		l.Fatal("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if IfExist(creds) {
		l.Info("user", creds.Login, "found, redirecting to /login")
		LoginHandler(w,r)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	_, _, err = db.DBClient.Collection("users").Add(r.Context(), map[string]interface{}{
		creds.Login:hashedPass,
	})
	if err != nil {
		l.Fatal("can't register new user with", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}
