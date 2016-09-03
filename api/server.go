package api

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

var router = mux.NewRouter()

func addRoutes() {
	http.ListenAndServe(":4153", router)
}
