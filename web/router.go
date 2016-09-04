package web

import (
	"database/sql"
	"encoding/json"
	"net/http"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
	"github.com/seikho/docktor/api"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/sqlite3" // reasons
)

var router = mux.NewRouter()

const path = "unix:///var/run/docker.sock"

// Start ...
func Start() {
	router.HandleFunc("/hosts", getHosts)
	router.HandleFunc("/containers", getContainers)
	router.HandleFunc("/imagews", getImages)
	router.HandleFunc("/stats/{id}", getStats)
	http.ListenAndServe(":8080", router)
}

func getHosts(res http.ResponseWriter, req *http.Request) {
	sql.Open("sqlite3", "docktor.sqlite")
	res.Write([]byte(":)"))
}

func getContainers(res http.ResponseWriter, req *http.Request) {
	containers := api.GetContainers(getClient())
	list, _ := json.Marshal(containers)
	res.Write(list)
}

func getImages(res http.ResponseWriter, req *http.Request) {
	images := api.GetImages(getClient())
	list, _ := json.Marshal(images)
	res.Write(list)
}

func getClient() *docker.Client {
	return api.GetClient(path)
}

func getStats(cres http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	containerID := vars["id"]
	client := getClient()
	stats := make(chan *docker.Stats)
	done := make(chan bool)
	opts := docker.StatsOptions{Stream: false, ID: containerID, Stats: stats, Done: done}

	go func() {
		err := client.Stats(opts)
		if err != nil {

		}
	}()
}

// IDObject ...
type IDObject interface {
	GetID() string
}

// Msg ...
type Msg struct {
	data    []string
	message string
}
