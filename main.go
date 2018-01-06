package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	r "github.com/GoRethink/gorethink"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

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

		all := make([]string, 0)

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

func main() {

	// load in config file from settings folder
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("missing config file; expecting a file named [%v] in current app directory\n", viper.ConfigFileUsed())
	}

	// validate mandatory config options
	if err := validateCfg(); err != nil {
		log.Fatalln(err)
	}

	// create database connection
	session, err := r.Connect(
		r.ConnectOpts{
			Address:  viper.GetString("re_ip") + ":" + viper.GetString("re_port"),
			Database: viper.GetString("re_database"),
		},
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "error connecting to rethinkDB"))
	}

	// pass session to a handler
	http.Handle("/playground", playgroundHandler(session))

	// listen and serve
	addr := viper.GetString("http_address") + ":" + viper.GetString("http_port")
	http.ListenAndServe(addr, nil)
}

func validateCfg() error {
	// mandatory config options, if these are missing program will crash
	config := []string{
		"re_database",
		"re_port",
		"re_ip",
		"http_address",
		"http_port",
	}

	for _, c := range config {
		if !viper.IsSet(c) {
			return errors.Errorf("Error. Missing mandatory config option: %v", c)
		}
	}

	return nil
}
