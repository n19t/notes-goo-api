package routes

import (
	"notes-goo-api/routes/tile"

	"github.com/gin-gonic/gin"
)

// Routes:
func Routes(r *gin.RouterGroup, dbs map[string]interface{}) {
	tile.Route(r.Group("tile"), dbs)
}
