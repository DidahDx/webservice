package main

import(
"net/http"
"github.com/DidahDx/webservice/controllers"
)


func main(){

controllers.RegisterControllers()
http.ListenAndServe(":3000",nil)

}