package config

import "google.golang.org/grpc"

func GetClientGRPCData() *grpc.ClientConn {
	return clientGRPCData
}
