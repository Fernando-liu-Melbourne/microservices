package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Fernando-liu-Melbourne/microservices/services"
	"github.com/Fernando-liu-Melbourne/microservices/utils"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Touched!!!x")
	var jsonValue []byte
	var applicationError *utils.ApplicationError

	userId, err :=strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		applicationError = &utils.ApplicationError{
			Message: "User ID must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "StatusBadRequest",
		}
		jsonValue, _ := json.Marshal(applicationError)
		response.Write(jsonValue)
		return
	}

	user, applicationError := services.GetUser(uint64(userId))
	if applicationError != nil {
		jsonValue, _ = json.Marshal(applicationError)
		response.WriteHeader(applicationError.StatusCode)
		response.Write(jsonValue)
		return
	}

	// return user to client
	jsonValue, _ = json.Marshal(user)
	response.Write(jsonValue)
}