package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lanpaiva/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
