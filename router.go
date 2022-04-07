package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vokestd/TestingSystem-server/endpoints"
)

func route(r *gin.Engine) {
	r.GET("/api/getQuestions", endpoints.GetQuestions)
	r.POST("/api/submitData", endpoints.SubmitData)
}
