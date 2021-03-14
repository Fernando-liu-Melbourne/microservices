package main

import (
	"fmt"
	app "github.com/Fernando-liu-Melbourne/microservices/app"
)

func main() {
	for i :=0; i<100 ;i++{
		fmt.Printf("ssss %d \r\n", i)
	}
	app.StartApp()


}