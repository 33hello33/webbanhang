package api

import (
	"net/http"
	"net/url"
	token "webbanhang/token"

	"github.com/gin-gonic/gin"
)

const authorizationPayloadKey = "authorization_payload"

func authMiddleware(tokenMaker token.PasetoMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			//err := errors.New("authorization header is not provided")
			//ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))

			location := url.URL{Path: "/login"}
			ctx.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		payload, err := tokenMaker.VerifyToken(cookie)
		if err != nil {
			if err == token.ErrExpireToken {
				location := url.URL{Path: "/token/renew"}
				ctx.Redirect(http.StatusFound, location.RequestURI())
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
				return
			}
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
