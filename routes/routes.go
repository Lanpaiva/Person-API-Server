package routes

import (
	"log"
	"net/http"

	"github.com/lanpaiva/go-rest-api/controller"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/api", controller.AllPersons)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
