package main

import (
	"fmt"
	"net"
	"os"
	"spos/queueing/amqp"
	"spos/queueing/auth"

	"github.com/s-pos/go-utils/config"
	authPb "github.com/s-pos/protobuf/go/auth"
	"google.golang.org/grpc"
)

func main() {
	log := config.Logrus()

	// amqp client channel
	producer := amqp.New()

	// grpc handler
	authGrpc := auth.NewAuthGRPC(producer)

	// run grpc server
	serv := grpc.NewServer()
	authPb.RegisterUserAuthServiceServer(serv, authGrpc)

	log.Fatal(runGrpc(serv))
}

func runGrpc(grpcServer *grpc.Server) error {
	grpcPort := fmt.Sprintf(":%s", os.Getenv("GRPC_PORT"))

	l, err := net.Listen("tcp", grpcPort)
	if err != nil {
		panic(err)
	}

	return grpcServer.Serve(l)
}
