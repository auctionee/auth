package main

import (
	"github.com/auctionee/auth/pkg/db"
	"github.com/auctionee/auth/pkg/server"
	"sync"
)
const DEFAULT_PORT = 8080
func main(){
	wg := sync.WaitGroup{}

	s := server.NewAuthServer(DEFAULT_PORT)
	db.CreateClient(s.Ctx)
	wg.Add(1)
	go func() {
		s.Start()
	}()
	wg.Wait()
}

