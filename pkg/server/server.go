package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simpleWebTemplate/config"
	_ "simpleWebTemplate/docs"
	"simpleWebTemplate/internal/service"
	"simpleWebTemplate/pkg/api"
	"simpleWebTemplate/pkg/dao/postgres"
	"syscall"
	"time"
)

type Server struct {
	BasePath string
	Port string
	Handler *gin.Engine
	Config *config.Config
	ApiRoute *api.API
}


func NewServer(config *config.Config) *Server {
	// initialize server, database access, handler

	service := &service.UserService{UserRepository: postgres.NewMockUserDAO()}

	server := Server{
		BasePath: "/",
		Port:     "8080",
		Handler:  NewRouter(),
		Config:   config,
		ApiRoute: &api.API{UserService: service},
	}
	return &server
}

func (server *Server) Serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s",server.Port),
		Handler: server.Handler,
	}

	// Register Web Service
	server.RegisterRouter()

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to listen " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}