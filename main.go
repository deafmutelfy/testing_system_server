package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vokestd/TestingSystem-server/logger"
	"github.com/vokestd/TestingSystem-server/model"
	"github.com/vokestd/TestingSystem-server/stuff"
)

func main() {
	model.TraitsGetInstance("./traits.yaml", &stuff.MarkupParserYaml{})
	model.QuestionsGetInstance("./questions.yaml", &stuff.MarkupParserYaml{})

	_ = godotenv.Load()

	logger.GetInstance().Infoln("Invoking Server Engine...")
	r := gin.Default()
	r.Use(cors.Default())

	route(r)

	port, isset := os.LookupEnv("PORT")
	if !isset {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
