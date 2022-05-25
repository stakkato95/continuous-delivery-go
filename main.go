package main

import (
	"fmt"
	"os"

	"github.com/stakkato95/continuous-delivery-go/app"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	fmt.Println("test change")
	a.Run(":8010")
}
