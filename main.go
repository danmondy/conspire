package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/danmondy/conspire/repo"
	"github.com/danmondy/conspire/routes"
)

var config Config

//init fire before main
func init() {
	//init config obj from file
	if err := readConfig(); err != nil {
		panic(err)
	}
}

func main() {
	//configure log file
	f, err := os.OpenFile("serverlog.txt", os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Log file opened for writing")
	log.Println(config.Db)

	db, err := sql.Open("sqlite3", config.Db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	routes.Repo = repo.NewRepo(db)

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server running on", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}

func readConfig() error {
	conf, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(conf, &config)
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Title, Version, Port, Db string
}
