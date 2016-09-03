package web

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/seikho/docktor/api"                 // reasons
	_ "gopkg.in/doug-martin/goqu.v3/adapters/sqlite3" // reasons
)

var router = mux.NewRouter()

func Start() {
	router.HandleFunc("/hosts", getHosts)
}

func getHosts(res http.ResponseWriter, req *http.Request) {
	sql.Open("sqlite3", "docktor.sqlite")
	res.Write([]byte(":)"))

}
