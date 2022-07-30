package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

type Server struct {
	e   *gin.Engine
	srv http.Server
}

func New() *Server {
	return &Server{
		e: gin.Default(),
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	return s.e
}

func (s *Server) Start() {
	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: s.SetupRoutes(),
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := s.srv.Close(); err != nil {
			log.Println("failed to shutdown server", err)
		}
	}()

	func() {
		if err := s.srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("server closed after interruption")
			} else {
				log.Println("unexpected server shutdown. err:", err)
			}
		}
	}()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
