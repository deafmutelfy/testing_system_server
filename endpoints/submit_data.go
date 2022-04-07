package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vokestd/TestingSystem-server/controller"
	"github.com/vokestd/TestingSystem-server/model"
)

type submitDataType struct {
	Auth struct {
		FirstName string
		LastName  string
	}
	Answers model.SubmitDataModel
}

func SubmitData(ctx *gin.Context) {
	dat := submitDataType{}
	ctx.BindJSON(&dat)
	fmt.Println(dat)

	controller.ProcessData(&dat.Answers)

	ctx.String(http.StatusOK, "")
}
