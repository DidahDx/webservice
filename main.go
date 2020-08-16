package main

import (
	"net/http"
	"os"

	"github.com/DidahDx/webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)

}
