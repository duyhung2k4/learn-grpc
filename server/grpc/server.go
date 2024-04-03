package grpc

import (
	"app/grpc/api"
	"app/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ServerGRPC() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalln(err)
	}

	creds, errCreds := credentials.NewServerTLSFromFile(
		"keys/public.pem",
		"keys/private.pem",
	)

	if errCreds != nil {
		log.Fatalln(errCreds)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}
	s := grpc.NewServer(opts...)
	proto.RegisterDataServiceServer(s, api.NewDataGRPC())

	log.Println("Start server grpc")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
