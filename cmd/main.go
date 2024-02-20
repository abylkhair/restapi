package cmd

import (
	"log"
	"restapi"
)

func main() {

	srv := new(restapi.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error")
	}
}
