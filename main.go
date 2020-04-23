package main

import(
	"fmt"
"github.com/DidahDx/webservice/models"
)


func main(){

u:=models.User{
	ID:2,
	FirstName:"John",
	SecondName:"Test",
   }

	fmt.Println("Hello",u)

}