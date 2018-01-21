package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/popwalker/studygo/grace/gracehttp"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"got ok",
		})
	})

	srv := &http.Server{
		Addr:"localhost:8080",
		Handler:r,
	}

	err := gracehttp.Serve(srv)
	if err != nil {
		panic(err)
	}

}
