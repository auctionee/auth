package handlers

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"github.com/auctionee/auth/pkg/data"
	"github.com/auctionee/auth/pkg/db"
	"github.com/auctionee/auth/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"io"
	"io/ioutil"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger(r.Context())
	creds := &data.Credentials{}
	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)
	err := json.NewDecoder(reader).Decode(creds)
	if err != nil {
		l.Fatal("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(b)
	if err := db.IfExist(r.Context(), creds); err != nil {
		l.Info("user", creds.Login, "found, redirecting to /login")
		LoginHandler(w,r)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	_, err = db.DBClient.Collection("users").Doc("users").Set(r.Context(), map[string]string{
		creds.Login:string(hashedPass),
	}, firestore.MergeAll)

	if err != nil {
		l.Fatal("can't register new user with", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}
