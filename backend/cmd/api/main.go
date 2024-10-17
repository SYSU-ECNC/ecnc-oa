package main

import "github.com/SYSU-ECNC/ecnc-oa/backend/internal/application"

func main() {
	app := application.NewApplication()
	app.InitApplication()
	app.Run()
}
