package webservice

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {

	router := NewRouter()
	http.Handle("/", router)

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error ocurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
