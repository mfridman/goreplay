package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	r "github.com/GoRethink/gorethink"
	"github.com/pkg/errors"
)

// The values below are specific to rethinkdb and should be adjusted for your environemnt
var (
	// DATABASE is the name of db connecting to, default is test
	DATABASE = "test"
	// IP is the location of rethinkdb
	IP = "localhost"
	// PORT, default rethinkdb port is 28015
	PORT = "28015"
)

// Listen and serve, i.e,. http://localhost:3001/playground
//
// Can use gin for live reloading without building binary each time
// go get -u github.com/codegangsta/gin
// verify it is installed by running: gin -h
// then from root of application run: gin run main.go
// and access the app from http://localhost:3000/playground
// NOT A TYPO. port 3000 is the default proxy port gin uses, our app is still listening/servering on port 3001
var (
	// HTTPHostName
	HTTPHostName = "localhost"
	// HTTPPort
	HTTPPort = "3001"
)

func main() {

	// create database connection
	session, err := r.Connect(
		r.ConnectOpts{
			Address:  IP + ":" + PORT,
			Database: DATABASE,
		},
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error connecting to rethinkDB"))
	}

	// pass session to a handler
	http.Handle("/playground", playgroundHandler(session))

	// listen and serve
	http.ListenAndServe(HTTPHostName+":"+HTTPPort, nil)
}

func playgroundHandler(session *r.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if !session.IsConnected() {
			http.Error(w, "rethinkdb session not available, check session.IsConnected()", http.StatusInternalServerError)
			return
		}
		/*
			start of play
		*/

		cur, err := r.TableList().Run(session)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		var all []string

		if err := cur.All(&all); err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		/*
			end of play
		*/

		// pretty print items from rethinkdb
		j, err := json.MarshalIndent(all, "", "\t")
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(j))
	})
}
