package routes

import (
	"log"
	"net/http"

	"go-api-rest/controllers"
	"go-api-rest/middleware"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonality).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindPersonalityById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", r))
}
