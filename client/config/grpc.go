package config

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func conectGRPCData() {
	var err error

	creds, errCreds := credentials.NewClientTLSFromFile("keys/public.pem", "localhost")
	if errCreds != nil {
		log.Fatalln(errCreds)
	}

	clientGRPCData, err = grpc.Dial("localhost:10000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
}
