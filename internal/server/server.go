package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/bank-transfer/internal/handler"
)

type Server struct {
	e   *gin.Engine
	srv http.Server
	h   *handler.Handler
}

func New(h *handler.Handler) *Server {
	return &Server{
		e: gin.Default(),
		h: h,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	s.e.GET("/healthCheck", s.h.HealthCheck)
	s.e.POST("/createAccount", s.h.CreateAccount)
	s.e.POST("/withdraw", s.h.WithdrawFromAccount)
	s.e.POST("/deposit", s.h.DepositToAccount)
	s.e.POST("/transfer", s.h.TransferToAccount)
	s.e.GET("/transaction", s.h.FindAllTransactions)
	s.e.GET("/transaction/filter", s.h.FindTransaction)

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
