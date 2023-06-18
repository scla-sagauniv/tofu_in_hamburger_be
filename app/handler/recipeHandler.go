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

type RecipeHandler struct {
	ingredientRainv1connect.UnimplementedRecipeServiceHandler
}

func (RecipeHandler) CreateRecipesByBatch(ctx context.Context, req *connect_go.Request[v1.CreateRecipesByBatchRequest]) (*connect_go.Response[v1.CreateRecipesByBatchResponse], error) {
	log.Println("Request headers: ", req.Header())
	var err error
	err = logic.DeleteAllRecipe()
	if err != nil {
		errMsg := err.Error()
		res := connect.NewResponse(&v1.CreateRecipesByBatchResponse{
			Error: &errMsg,
		})
		return res, nil
	}
	err = logic.BulkInsertToRecipe(req.Msg.Recipes)
	if err != nil {
		errMsg := err.Error()
		res := connect.NewResponse(&v1.CreateRecipesByBatchResponse{
			Error: &errMsg,
		})
		return res, nil
	}

	res := connect.NewResponse(&v1.CreateRecipesByBatchResponse{
		Error: nil,
	})

	return res, nil
}

func SearchRecipesByIngredients(ctx context.Context, req *connect_go.Request[v1.SearchRecipesByIngredientsRequest]) (*connect_go.Response[v1.SearchRecipesByIngredientResponse], error) {
	log.Println("Request headers: ", req.Header())
	recipes, err := logic.SelectRecipe(req.Msg.Ingredients)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	res := connect.NewResponse(&v1.SearchRecipesByIngredientResponse{
		Recipes: recipes,
	})
	return res, nil
}
