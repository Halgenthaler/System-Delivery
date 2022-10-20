package main

import (
	"fmt"
	"log"

	"github.com/Halgenthaler/System-Delivery/app/route"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error .env file")
	}
}

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[2:8])

}
