package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var HOST, PORT, ADDR string
var parent context.Context
var signalling chan os.Signal
var logger *log.Logger
var server *http.Server

func main() {
	go func() {
		logger.Printf("listening for HTTP(S) at %s", ADDR)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()
	<-signalling
	logger.Printf("shutting down HTTP(S) server listening at %s", ADDR)
	shutCtx, cancel := context.WithTimeout(parent, time.Second*5)
	defer cancel()
	err := server.Shutdown(shutCtx)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	logger.Printf("HTTP(S) server listening at %s gracefully shut down", ADDR)
}

func init() {
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
	ADDR = HOST + ":" + PORT

	parent = context.Background()

	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signalling = make(chan os.Signal, len(signals))
	signal.Notify(signalling, signals...)

	HANDLER := &NullHandler{}

	logger = log.Default()

	server = &http.Server{
		Addr:        ADDR,
		Handler:     HANDLER,
		BaseContext: func(l net.Listener) context.Context { return parent },
	}

}

type NullHandler struct{}

func (h *NullHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
