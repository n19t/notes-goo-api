package tile

import "github.com/gin-gonic/gin"

type handler struct {
	dbs map[string]interface{}
}

func NewHandler(dbs map[string]interface{}) *handler {
	return &handler{
		dbs: dbs,
	}
}

// Route :
func Route(r *gin.RouterGroup, dbs map[string]interface{}) {

	h := NewHandler(dbs)

	r.GET("", h.fetchAll)
}
