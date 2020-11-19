package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	router.GET("/health", api.HealthCheck)

	go func() {
		err = router.Run(net.JoinHostPort("", apiPort))
		log.Println("Starting api Server ...")
		if err != nil {
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

	// TODO: shutdown(router)

	log.Println("Server exiting")
}
