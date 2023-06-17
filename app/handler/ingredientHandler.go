package handler

import (
	"context"
	"log"
	v1 "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1"
	"tofu_in_hamburger_be/gen/rpc/ingredientRain/v1/ingredientRainv1connect"
	"tofu_in_hamburger_be/logic"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
)

type IngredientHandler struct {
	ingredientRainv1connect.UnimplementedIngredientServiceHandler
}

var Streams []*connect_go.ServerStream[v1.StreamIngredientResponse]

func (IngredientHandler) GetIngredientList(ctx context.Context, req *connect_go.Request[v1.GetIngredientListRequest]) (*connect_go.Response[v1.GetIngredientListResponse], error) {
	log.Println("Request headers: ", req.Header())
	ingredients, err := logic.SelectAllIngredient()
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&v1.GetIngredientListResponse{
		Ingredients: ingredients,
	})
	return res, nil
}

func (IngredientHandler) StreamIngredient(ctx context.Context, req *connect_go.Request[v1.StreamIngredientRequest], stm *connect_go.ServerStream[v1.StreamIngredientResponse]) error {
	log.Println("Request headers: ", req.Header())
	Streams = append(Streams, stm)
	log.Println("Stream start(len(Streams): ", len(Streams), ")")
	for {
		// セッションを切らさないための無限ループ
	}
}

func (IngredientHandler) SendIngredients(ctx context.Context, req *connect_go.Request[v1.SendIngredientsRequst]) (*connect_go.Response[v1.SendIngredientsResponse], error) {
	log.Println("Request headers: ", req.Header())

	logic.RecievedIngredients = append(logic.RecievedIngredients, req.Msg.Ingredients...)

	for _, stream := range Streams {
		selected := logic.IngredientSelector()
		stream.Send(&v1.StreamIngredientResponse{
			Ingredients: selected,
		})
	}

	res := connect.NewResponse(&v1.SendIngredientsResponse{})
	return res, nil
}
