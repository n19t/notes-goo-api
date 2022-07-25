package routes

import (
	"log"
	"notes-goo-api/internal/config"
	"notes-goo-api/internal/db"
	"notes-goo-api/internal/fbauth"
	_routes "notes-goo-api/routes"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	client, err := fbauth.InitAuth()
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(Cors())
	r.Use(fbauth.AuthJWT(client))

	c, err := config.ReadConfig()
	if err == nil {
		dbs, _ := db.InitDb(c)
		_routes.Routes(r.Group(""), dbs)
		r.Run(c.Port)
	} else {
		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}
