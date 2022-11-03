package main

import (
	"github.com/emirkosuta/golang-auth/controller"
	"github.com/emirkosuta/golang-auth/model"
)

func Main() {
	model.Init()
	controller.Start()
}
