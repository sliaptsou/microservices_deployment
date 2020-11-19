package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/sliaptsou/backend/handler"
	"github.com/sliaptsou/backend/proto"
)

const (
	svcHost = "SVC_HOST"
	svcPort = "SVC_PORT"

	maxMessageSize       = 1024 * 1024 * 50
	maxConcurrentStreams = 300
)

func main() {
	serviceHost, ok := os.LookupEnv(svcHost)
	if !ok {
		log.Fatalf("%s env variable is not set", svcHost)
	}

	servicePort, ok := os.LookupEnv(svcPort)
	if !ok {
		log.Fatalf("%s env variable is not set", svcPort)
	}

	lis, err := net.Listen("tcp", net.JoinHostPort(serviceHost, servicePort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(maxMessageSize),
		grpc.MaxSendMsgSize(maxMessageSize),
		grpc.MaxConcurrentStreams(maxConcurrentStreams),
	)

	proto.RegisterBackendServer(grpcServer, handler.NewBackendServer())

	go func() {
		log.Println("Starting service Server ...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Server ...")

	// TODO: shutdown(router)

	log.Println("Server exiting")
}
