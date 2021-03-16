package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Fernando-liu-Melbourne/microservices/services"
	"log"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Touched!!!x")
	userId, err :=strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("user id must be a number"))
		return
	}
	log.Println("User ID is", userId)

	user, err := services.GetUser(uint64(userId))
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}