package webservice

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func StartWebServer(port string) {

	router := NewRouter()
	http.Handle("/", router)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                            // All origins
		AllowedMethods: []string{"GET", "DELETE", "POST", "PUT"}, // Allowing only get, just an example
	})

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, c.Handler(router))

	if err != nil {
		log.Println("An error ocurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
