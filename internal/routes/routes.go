package routes

import (
	"net/http"
	"notes-goo-api/internal/config"
	"notes-goo-api/internal/db"
	_routes "notes-goo-api/routes"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	c, _ := config.ReadConfig("")

	dbs, _ := db.InitDb(c)

	_routes.Routes(r.Group(""), dbs)

	r.Run(c.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
