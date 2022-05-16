package main

import (
	"os"

	"github.com/stakkato95/continuous-delivery-go/app"
)

func main() {
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
