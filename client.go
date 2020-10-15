package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/go-recipe-grpc/recipe"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	r := recipe.NewRecipeServiceClient(conn)

	response, err := r.GetRecipe(context.Background(), &recipe.RecipeRequest{MainIngredient: "Lamb"})
	if err != nil {
		log.Fatalf("Error when calling GetRecipe: %s", err)
	}
	log.Printf("Response from server: %s", response.RecipeName)
}
