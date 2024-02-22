package cmd

import (
	"log"
	"restapi"
	"restapi/pkg/handler"
	"restapi/pkg/repository"
	"restapi/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapi.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error")
	}
}
