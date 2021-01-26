package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sliaptsou/backend/proto"
	"google.golang.org/grpc"
)

const (
	apiPort = "API_PORT"
	svcHost = "SVC_HOST"
	svcPort = "SVC_PORT"

	maxMessageSize = 1024 * 1024 * 50
)

func main() {
	apiPort, ok := os.LookupEnv(apiPort)
	if !ok {
		log.Fatalf("%s env variable is not set", apiPort)
	}

	serviceHost, ok := os.LookupEnv(svcHost)
	if !ok {
		log.Fatalf("%s env variable is not set", svcHost)
	}

	servicePort, ok := os.LookupEnv(svcPort)
	if !ok {
		log.Fatalf("%s env variable is not set", svcPort)
	}

	var conn *grpc.ClientConn
	conn, err := grpc.DialContext(context.Background(), net.JoinHostPort(serviceHost, servicePort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMessageSize),
			grpc.MaxCallSendMsgSize(maxMessageSize),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	client := proto.NewBackendClient(conn)
	defer conn.Close()

	api := NewApi(client)
	router := gin.Default()
	router.GET("/count", api.GetQueryCount)
	router.GET("/healthz", api.HealthCheck)

	router.GET("/entity", api.GetList)
	router.GET("/entity/:id", api.GetOne)
	router.POST("/entity", api.Create)
	router.PUT("/entity/:id",api.Update)
	router.DELETE("/entity/:id", api.Delete)
	srv := &http.Server{
		Addr:    net.JoinHostPort("", apiPort),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		err := srv.ListenAndServe()
		log.Println("Starting api Server ...")
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error while running server %v", err)
			return
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	signals := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("Shutting down Server ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	os.Exit(0)
}
