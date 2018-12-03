package main

import (
	"github.com/dasetova/truora_test/utils"
	"github.com/dasetova/truora_test/webservice"
)

func main() {
	utils.MigrateDB()
	webservice.StartWebServer("6767")
}
