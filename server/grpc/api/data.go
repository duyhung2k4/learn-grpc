package api

import (
	"app/grpc/proto"
	"context"
	"io"
	"time"
)

type dataGRPC struct {
	proto.UnsafeDataServiceServer
}

func (g *dataGRPC) SendTextOneToOne(ctx context.Context, req *proto.DataReq) (*proto.DataRes, error) {
	newText := req.Text + " - New"
	return &proto.DataRes{
		Mess: newText,
	}, nil
}

func (g *dataGRPC) SendTextOneToMany(req *proto.DataReq, stream proto.DataService_SendTextOneToManyServer) error {
	text := req.Text
	for _, c := range text {
		if err := stream.Send(&proto.DataRes{Mess: string(c)}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func (g *dataGRPC) SendTextManyToOne(stream proto.DataService_SendTextManyToOneServer) error {
	longText := ""
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}

		newText := result.Text + "-new"
		longText += newText + " "
	}

	stream.SendAndClose(&proto.DataRes{
		Mess: longText,
	})
	return nil
}

func (g *dataGRPC) SendTextManyToMany(stream proto.DataService_SendTextManyToManyServer) error {
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if result.Text == "0" {
			stream.Context().Done()
			break
		}

		text := result.Text + "-new"
		stream.Send(&proto.DataRes{
			Mess: text,
		})
	}
	return nil
}

func NewDataGRPC() proto.DataServiceServer {
	return &dataGRPC{}
}
