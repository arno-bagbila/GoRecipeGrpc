package recipe

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetRecipe(cyx context.Context, in *RecipeRequest) (*RecipeReply, error) {
	log.Printf("Received message from client: %s", in.MainIngredient)
	if in.MainIngredient == "Chicken" {
		return &RecipeReply{RecipeName: "Chicken Masala"}, nil
	} else {
		return &RecipeReply{RecipeName: "Lamb Masala"}, nil
	}

}
