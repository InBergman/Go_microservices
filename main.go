package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	// router.GET("/", func(context *gin.Context) {

	// 	context.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		GIN.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })
	initializeRoutes()
	router.Run()
}
