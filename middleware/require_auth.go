package middleware

import (
	"fmt"
	"go-jwt/initializers"
	"go-jwt/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func MiddleWare(ctx *gin.Context) {
	// get the cookie off req
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	// decode validate it

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header)
		}
		return []byte("dasdasihdgashdjkahskhas"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the app
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		// find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		// attach to req
		ctx.Set("user", user)
		// continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
