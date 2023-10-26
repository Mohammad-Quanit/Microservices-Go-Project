package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (l *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good Bye"))
}
