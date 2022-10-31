package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewTokenResponse struct {
	Token          string    `json:"token"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
}

func (server *Server) renewTokenHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "renew_token.html", nil)
}

func (server *Server) renewToken(ctx *gin.Context) {
	var req renewTokenRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errResponse(err))
		return
	}

	if session.IsBlocked == true {
		err := fmt.Errorf("Block session")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("Incorrect session user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	if time.Now().After(session.ExpiredAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	token, payload, err := server.tokenMaker.CreateToken(refreshPayload.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	res := renewTokenResponse{
		Token:          token,
		TokenExpiredAt: payload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, res)
}
