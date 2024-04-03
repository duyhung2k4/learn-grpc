package controller

import (
	"app/config"
	"app/grpc/proto"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/render"
)

type dataController struct {
	clientGRPCData proto.DataServiceClient
}

type DataController interface {
	SendTextOneToOne(w http.ResponseWriter, r *http.Request)
	SendTextOneToMany(w http.ResponseWriter, r *http.Request)
	SendTextManyToOne(w http.ResponseWriter, r *http.Request)
	SendTextManyToMany(w http.ResponseWriter, r *http.Request)
}

type Mess struct {
	Text string `json:"text"`
}

func (c *dataController) SendTextOneToOne(w http.ResponseWriter, r *http.Request) {
	var mess Mess
	json.NewDecoder(r.Body).Decode(&mess)

	result, err := c.clientGRPCData.SendTextOneToOne(context.Background(), &proto.DataReq{
		Text: mess.Text,
	})

	if err != nil {
		log.Println("Error: ", err)
		render.JSON(w, r, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"mess": result.Mess,
	})
}

func (s *dataController) SendTextOneToMany(w http.ResponseWriter, r *http.Request) {
	var mess Mess
	json.NewDecoder(r.Body).Decode(&mess)

	result, err := s.clientGRPCData.SendTextOneToMany(r.Context(), &proto.DataReq{Text: mess.Text})
	if err != nil {
		render.JSON(w, r, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	list := []map[string]interface{}{}

	for {
		char, err := result.Recv()
		if err == io.EOF {
			break
		}

		list = append(list, map[string]interface{}{
			"char": char.Mess,
		})
	}

	render.JSON(w, r, map[string]interface{}{
		"list": list,
	})
}

func (c *dataController) SendTextManyToOne(w http.ResponseWriter, r *http.Request) {
	var mess Mess
	json.NewDecoder(r.Body).Decode(&mess)

	stream, err := c.clientGRPCData.SendTextManyToOne(context.Background())
	if err != nil {
		render.JSON(w, r, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	for _, c := range mess.Text {
		errSend := stream.Send(&proto.DataReq{
			Text: string(c),
		})
		if errSend != nil {
			render.JSON(w, r, map[string]interface{}{
				"error": errSend.Error(),
			})
			return
		}
		time.Sleep(time.Second * 1)
	}

	res, errRes := stream.CloseAndRecv()
	if errRes != nil {
		render.JSON(w, r, map[string]interface{}{
			"error": errRes.Error(),
		})
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"mess": res.Mess,
	})
}

func (c *dataController) SendTextManyToMany(w http.ResponseWriter, r *http.Request) {
	var mess Mess
	json.NewDecoder(r.Body).Decode(&mess)

	stream, err := c.clientGRPCData.SendTextManyToMany(context.Background())
	if err != nil {
		render.JSON(w, r, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	text := ""
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for _, c := range mess.Text {
			errRes := stream.Send(&proto.DataReq{
				Text: string(c),
			})

			if errRes != nil {
				render.JSON(w, r, map[string]interface{}{
					"error": errRes.Error(),
				})
				return
			}
		}
		wg.Done()
		stream.CloseSend()
	}()

	go func() {
		for {
			res, errRes := stream.Recv()
			if errRes == io.EOF {
				break
			}

			text += res.Mess + " "
		}
		wg.Done()
	}()

	wg.Wait()

	render.JSON(w, r, map[string]interface{}{
		"text": text,
	})

}

func NewDataController() DataController {
	return &dataController{
		clientGRPCData: proto.NewDataServiceClient(config.GetClientGRPCData()),
	}
}
