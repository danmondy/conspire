package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

//We need to open a connection to a log file and a db in the main method so they will remain
//open until the application closes. Not sure if there is a way around this.
func main() {
	//configure log file
	f, err := os.OpenFile("serverlog.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	log.SetOutput(f)
	log.Println("logfile: ", f.Name)

	//configure db
	db, err := sql.Open("sqlite3", config.Db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//register routes, server listen.
	routes.Repo = repo.NewRepo(db)
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Println("Server running on", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, r))


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
