package main

import (
	"log"
	"os"

	"github.com/dasetova/truora_test/utils"
	"github.com/dasetova/truora_test/webservice"
)

func main() {
	err := utils.ValidateEnvVars()
	if err != nil {
		log.Fatal(err)
	}

	utils.MigrateDB()
	webservice.StartWebServer(os.Getenv("API_PORT"))
}
