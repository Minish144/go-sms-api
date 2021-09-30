package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run() {
	router := mux.NewRouter()

	http.Handle("/", router)

	connectionString := getListeningAddress()
	logrus.Infof("Server is now listening on http://%s", connectionString)
	http.ListenAndServe(connectionString, nil)
}

func getListeningAddress() string {
	host := viper.GetString("server.interface")
	port := viper.GetString("server.port")
	return fmt.Sprintf("%s:%s", host, port)
}
