package main

import (
	"goservice"
)

func main() {
	goservice.NewService(8080).Run()
}
