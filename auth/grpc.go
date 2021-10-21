package auth

import (
	"context"
	"fmt"
	"spos/queueing/amqp"

	pb "github.com/s-pos/protobuf/go/auth"
)

type authGrpc struct {
	pb.UnimplementedUserAuthServiceServer
	publish amqp.Producer
}

func NewAuthGRPC(p amqp.Producer) pb.UserAuthServiceServer {
	return &authGrpc{
		publish: p,
	}
}

func (a *authGrpc) SendEmailVerification(ctx context.Context, req *pb.Verification) (*pb.VerificationReply, error) {
	var (
		res = new(pb.VerificationReply)
		err error
	)

	err = a.publish.PublishMessage(amqp.RegisterVerification, req)
	if err != nil {
		return nil, err
	}

	res = &pb.VerificationReply{
		Message: fmt.Sprintf("Email berhasil dikirim ke %s", req.GetEmail()),
	}

	return res, err
}
