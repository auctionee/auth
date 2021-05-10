package handlers

import (
	"encoding/json"
	"github.com/auctionee/auth/pkg/data"
	"github.com/auctionee/auth/pkg/db"
	"github.com/auctionee/auth/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	l := logger.GetLogger(r.Context())
	creds := &data.Credentials{}
	res := &data.AuthResponse{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		l.Fatal("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	passHash := db.GetHashedPassword(r.Context(), creds)
	if passHash == "" {
		l.Info("user", creds.Login, "not found")
		w.WriteHeader(http.StatusBadRequest)
		_, errWrite := w.Write(res.ToJSON("false", "user not found"))
		if errWrite != nil{
			l.Info("error writing response in login handler: ", errWrite)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(creds.Password))
	if err != nil{
		l.Info("wrong password for user ",creds.Login)
		_, errWrite := w.Write(res.ToJSON("false", "wrong password"))
		if errWrite != nil{
			l.Info("error writing response in login handler: ", errWrite)
		}
		return
	}

	_, errWrite := w.Write(res.ToJSON("true", ""))
	if errWrite != nil{
		l.Info("error writing response in login handler: ", errWrite)
	}
}