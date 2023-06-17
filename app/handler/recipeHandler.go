package handler

import (
	"context"
	"errors"
	"log"

	v1 "tofu_in_hamburger_be/gen/rpc/ingredientRain/v1"
	"tofu_in_hamburger_be/gen/rpc/ingredientRain/v1/ingredientRainv1connect"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
)

type RecipeHandler struct {
	ingredientRainv1connect.UnimplementedRecipeServiceHandler
}

func (RecipeHandler) CreateRecipesByBatch(ctx context.Context, req *connect_go.Request[v1.CreateRecipesByBatchRequest]) (*connect_go.Response[v1.CreateRecipesByBatchResponse], error) {
	log.Println("Request headers: ", req)
	err := errors.New("error")
	res := connect.NewResponse(&v1.CreateRecipesByBatchResponse{
		Error: err.Error(),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
