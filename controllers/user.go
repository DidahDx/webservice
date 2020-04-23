package controllers

import(
	"net/http"
	"regexp"
)

type userController struct{
	userIDPattern *regexp.Regexp
}

//binding a function into a type
func (uc userController) ServeHTTP(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello user controller"))
}

func newUserController() *userController{
	return &userController{
		userIDPattern: regexp.MustCompile(`^/user/(\d+)/`),
	}
}