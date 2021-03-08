package main

import (
	"main/controllers"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	_ "github.com/joho/godotenv/autoload"
)

var linkController controllers.LinkController

func setup() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})
	r.GET("/:urlID", linkController.Match)
	r.POST("/create", linkController.Create)

	r.Run()
}

func main() {
	db.Init()

	setup()
}
