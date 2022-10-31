package api

import (
	"errors"
	"net/http"
	"net/url"
	"webbanhang/token"

	"github.com/gin-gonic/gin"
)

const authorizationPayloadKey = "authorization_payload"

func authMiddleware(tokenMaker token.PasetoMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")

		if err != nil {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		payload, err := tokenMaker.VerifyToken(cookie)
		if err != nil {
			location := url.URL{Path: "/token/renew"}
			ctx.Redirect(http.StatusFound, location.RequestURI())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
