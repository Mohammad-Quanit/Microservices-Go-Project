package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye  {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHttp(rw http.ResponseWriter, r *http.Request)  {
	// g.l.Println("Hello from Good Bye logger")
	rw.Write([]byte("Hello from Good Bye logger"))
}