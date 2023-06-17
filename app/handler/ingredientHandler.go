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
