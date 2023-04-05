package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"live-auth/internal/verify"
)


type Auth struct {
	Key  string    `form:"key"`	
}

func main() {

	
	router := gin.Default()

	auth := func (ctx *gin.Context) {	

		var auth Auth

		if ctx.ShouldBind(&auth) != nil {
			ctx.String(http.StatusUnauthorized, "unauthorized, my friend")
		}

		v := verify.New(auth.Key)

		if len(auth.Key) > 1 && v.Ok() {
			ctx.String(http.StatusOK, "wellcome, dude")
		}

		ctx.String(http.StatusUnauthorized, "unauthorized, my friend")
		return
	}

	router.POST("/auth", auth)

	router.Run(":8000")
}