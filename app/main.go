package main

import (
	controller "github.com/emirkosuta/golang-auth/controller"
	"github.com/emirkosuta/golang-auth/model"
)

func main() {
	model.Init()
	controller.Start()
}
