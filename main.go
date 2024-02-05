package main

import (
	"crud-api/app"
	"crud-api/helper"
	"crud-api/utils"
	"fmt"
)

func main() {
	config, err := utils.LoadConfig(".")
	helper.PanicIfError(err)

	err = app.Serve(config)
	helper.PanicIfError(err)

	fmt.Println("Server started on port", config.ServerAddress)
}
