package app

import (
	controller "github.com/Fernando-liu-Melbourne/microservices/controllers"
	"log"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", controller.GetUser)
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}


}
