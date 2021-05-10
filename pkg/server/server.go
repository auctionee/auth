package server

import (
	"context"
	"fmt"
	"github.com/auctionee/auth/pkg/handlers"
	"github.com/auctionee/auth/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type AuthServer struct {
	Ctx context.Context
	port string
	connectionTimeout time.Duration
}

func NewAuthServer(port int) *AuthServer  {
	ctx := logger.NewCtxWithLogger()
	return &AuthServer{
		port: ":"+strconv.Itoa(port),
		Ctx : ctx ,
	}
}
func (s *AuthServer) Start(){
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, GCloud!")
	})
	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.Handle("/register/", handlers.LoggerMiddleware(s.Ctx, http.HandlerFunc(handlers.RegisterHandler)))
	postRouter.Handle("/login/", handlers.LoggerMiddleware(s.Ctx, http.HandlerFunc(handlers.LoginHandler)))
	if err := http.ListenAndServe(s.port, r); err != nil{
		logger.GetLogger(s.Ctx).Fatal()
	}

}
