package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	_ "simpleWebTemplate/docs"
	"simpleWebTemplate/pkg/api"
	"syscall"
	"time"
)

type Server struct {
	BasePath string
	Port string
	Handler http.Handler
}


func NewServer() *Server{
	router := gin.Default()
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := router.Group("/v1")
	api.AddUserRoutes(v1)

	server := Server{
		BasePath: "/",
		Port: "8080",
		Handler: router,
	}
	return &server
}

func (server *Server) Serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s",server.Port),
		Handler: server.Handler,
	}

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