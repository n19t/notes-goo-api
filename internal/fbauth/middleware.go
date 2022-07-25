package fbauth

import (
	"log"
	"net/http"
	"strings"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	valName             = "FIREBASE_ID_TOKEN"
)

func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		authHeader := ctx.Request.Header.Get(authorizationHeader)
		// log.Println("authHeader", authHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := client.VerifyIDToken(ctx, token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
		}

		log.Println("Auth time:", time.Since(startTime))

		ctx.Set(valName, idToken)
		ctx.Next()
	}
}
