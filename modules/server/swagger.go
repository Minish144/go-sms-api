package server

import "net/http"

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "gen/swagger/api.swagger.json")
}
