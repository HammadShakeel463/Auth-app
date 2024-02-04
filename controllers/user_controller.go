package controllers

import (
	"go-jwt/initializers"
	"go-jwt/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	// get the email/pass off the req body
	var Body struct {
		Email    string
		Password string
	}
	if ctx.Bind(&Body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(Body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
	}

	// Create the user
	user := models.User{Email: Body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
	}
	// respond
	ctx.JSON(http.StatusOK, gin.H{})

}

func Login(ctx *gin.Context) {
	// get the email and pass off req body
	var Body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&Body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	// look up requested user

	var user = models.User{}
	initializers.DB.First(&user, "email=?", Body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	// compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or pass",
		})
		return
	}

	// generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("dasdasihdgashdjkahskhas"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "token error",
		})
		return
	}

	// send it back
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
