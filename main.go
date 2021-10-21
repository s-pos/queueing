package main

import (
	"fmt"
	"net"
	"os"
	"spos/queueing/amqp"
	"spos/queueing/auth"

	"github.com/s-pos/go-utils/config"
	authPb "github.com/s-pos/protobuf/go/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log := config.Logrus()

	// amqp client channel
	producer := amqp.New(log)

	// grpc handler
	authGrpc := auth.NewAuthGRPC(producer, log)

	// run grpc server
	serv := grpc.NewServer()
	authPb.RegisterUserAuthServiceServer(serv, authGrpc)

	log.Fatal(runGrpc(serv, log))
}

func runGrpc(grpcServer *grpc.Server, log *logrus.Logger) error {
	grpcPort := fmt.Sprintf(":%s", os.Getenv("GRPC_PORT"))

	l, err := net.Listen("tcp", grpcPort)
	if err != nil {
		panic(err)
	}

	log.Printf("GRPC Running on port %s", grpcPort)
	return grpcServer.Serve(l)
}
