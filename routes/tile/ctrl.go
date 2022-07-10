package tile

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *handler) fetchAll(c *gin.Context) {

	var tile []tile

	// Declare Context type object for managing multiple API requests
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, _ := h.dbs["bitkub"].(*mongo.Client)
	col := client.Database("").Collection("")
	cur, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cur.All(ctx, &tile); err != nil {
		panic(err)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, tile)
	}

}
