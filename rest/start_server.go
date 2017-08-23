package rest

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/jhidalgo3/microservice-go/dao"

	"github.com/spf13/viper"
)

const (
	defaultConfigFilePath  = "./_conf"
	configFilePathUsage    = "config file directory. Config file must be named 'conf_{env}.yml'."
	configFilePathFlagName = "configFilePath"
	envUsage               = "environment for app, prod, dev, test"
	envDefault             = "prod"
	envFlagname            = "env"
)

var configFilePath string
var env string

func init() {
	flag.StringVar(&configFilePath, configFilePathFlagName, defaultConfigFilePath, configFilePathUsage)
	flag.StringVar(&env, envFlagname, envDefault, envUsage)
	flag.Parse()
	configuration(configFilePath, env)
	setupDB(viper.GetString("database.URL"))
}

// StartServer starts server with REST handlers and initialise db connection pool
func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/banks/", commonHeaders(getBanksHandler)).Methods("GET")
	r.HandleFunc("/api/banks/{id:[0-9]+}", commonHeaders(getBankByIDHandler)).Methods("GET")
	r.HandleFunc("/api/banks/", commonHeaders(createBankHanlder)).Methods("POST")
	r.HandleFunc("/api/banks/{id:[0-9]+}", commonHeaders(deleteBankByIDHandler)).Methods("DELETE")
	r.HandleFunc("/api/banks/{id:[0-9]+}", commonHeaders(updateBankHanlder)).Methods("PUT")
	r.HandleFunc("/api/banks/", commonHeaders(deleteAllBanksHandler)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(viper.GetString("server.port"), r))
}

func configuration(path string, env string) {
	if isTest, _ := regexp.MatchString("/_test/", os.Args[0]); isTest {
		env = "test"
		path = "../_conf"
	}
	log.Println("Environment is: " + env + " configFilePath is: " + path)
	viper.SetConfigName("conf_" + env)
	viper.AddConfigPath(path) // working directory
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
	}
}

func setupDB(dbURL string) {
	var db, err = dao.NewBankAPI(dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
	}
	dao.DBAccess = db
}
