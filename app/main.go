package main

import (
	"log"
	"net/http"
	"tofu_in_hamburger_be/db"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/microsoft/go-mssqldb"

	"tofu_in_hamburger_be/gen/rpc/ingredientRain/v1/ingredientRainv1connect"
	myHandler "tofu_in_hamburger_be/handler"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	db.Init()
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"rpc.ingredientRain.v1.IngredientService", // 作成したサービスを指定
		"rpc.ingredientRain.v1.RecipeService",     // 作成したサービスを指定
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	path, handler := ingredientRainv1connect.NewIngredientServiceHandler(myHandler.IngredientHandler{})
	mux.Handle(path, handler)
	path, handler = ingredientRainv1connect.NewRecipeServiceHandler(myHandler.RecipeHandler{})
	mux.Handle(path, handler)
	log.Println("server is launched")
	http.ListenAndServe(
		"0.0.0.0:8080",
		cors.AllowAll().Handler(
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
	defer db.Db.Close()
}
