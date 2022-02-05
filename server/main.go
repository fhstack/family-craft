package main

import (
	"log"

	"github.com/fhstack/family-craft/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/family/cat/lockspin", handler.CatLockSpinHandler)

	log.Println("server start")
	r.Run()
}
