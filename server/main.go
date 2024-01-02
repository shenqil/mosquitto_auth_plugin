package main

import (
	"context"
	"log"
	"mosquitto_auth_plugin/mosq_err"
	"net"

	"google.golang.org/grpc"

	pb "mosquitto_auth_plugin/mosquitto_auth"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) BasicAuth(ctx context.Context, in *pb.BasicAuthRequest) (*pb.BasicAuthReply, error) {

	log.Printf("server listening at %+v", in)
	return &pb.BasicAuthReply{Code: mosq_err.MOSQ_ERR_SUCCESS}, nil
}

func (s *server) AclCheck(ctx context.Context, in *pb.AclCheckRequest) (*pb.AclCheckReply, error) {
	log.Printf("server listening at %+v", in)
	return &pb.AclCheckReply{Code: mosq_err.MOSQ_ERR_SUCCESS}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":10086")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
