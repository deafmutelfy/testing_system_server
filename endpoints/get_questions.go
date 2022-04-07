package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/vokestd/TestingSystem-server/model"
	"net/http"
)

func GetQuestions(ctx *gin.Context) {
	q := model.QuestionsGetInstance("", nil)

	ctx.JSON(http.StatusOK, *q)
}
