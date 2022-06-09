package main

import (
	"fmt"
	"os"

	"github.com/stakkato95/continuous-delivery-go/app"
)

func main() {
	a := app.App{}

	fmt.Printf("APP_DB_USERNAME: %s\n", os.Getenv("APP_DB_USERNAME"))
	fmt.Printf("APP_DB_PASSWORD: %s\n", os.Getenv("APP_DB_PASSWORD"))
	fmt.Printf("APP_DB_NAME: %s\n", os.Getenv("APP_DB_NAME"))

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	fmt.Println("server is listening...")
	a.Run(":8010")
}
