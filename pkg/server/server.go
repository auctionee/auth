package server

import (
	"context"
	"github.com/auctionee/auth/pkg/handlers"
	"github.com/auctionee/auth/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type AuthServer struct {
	ctx context.Context
	port string
	connectionTimeout time.Duration
}

func NewAuthServer(port int) *AuthServer  {

	return &AuthServer{
		port: ":"+strconv.Itoa(port),
		ctx : logger.NewCtxWithLogger(),
	}
}
func (s *AuthServer) Start(){
	r := mux.NewRouter()
	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.Handle("/register/", handlers.LoggerMiddleware(s.ctx, http.HandlerFunc(handlers.RegisterHandler)))
	postRouter.Handle("/login/", handlers.LoggerMiddleware(s.ctx, http.HandlerFunc(handlers.LoginHandler)))
	if err := http.ListenAndServe(s.port, r); err != nil{
		logger.GetLogger(s.ctx).Fatal()
	}

}
