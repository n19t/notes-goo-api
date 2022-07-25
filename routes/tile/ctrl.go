package tile

import (
	"context"
	"log"
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

	client, _ := h.dbs["notes-goo"].(*mongo.Database)
	col := client.Collection("tile")
	cur, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Panic(err)
	}

	if err = cur.All(ctx, &tile); err != nil {
		log.Panic(err)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, tile)
	}

}
