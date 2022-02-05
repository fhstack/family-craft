package main

import (
	"log"

	"github.com/fhstack/family-craft/server/handler"
	"github.com/fhstack/family-craft/server/hardware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/family/cat/lockspin/:angle", handler.CatLockSpin)

	hardware.Init()
	log.Println("server start")
	r.Run()
}
