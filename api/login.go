package api

import (
	"database/sql"
	"net/http"
	"time"

	db "webbanhang/db/sqlc"
	"webbanhang/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) loginHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
}

type loginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type loginResponse struct {
	SessionID             uuid.UUID `json:"session_id"`
	Token                 string    `json:"token"`
	TokenExpiredAt        time.Time `json:"token_expired_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, account.HashPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	token, payload, err := server.tokenMaker.CreateToken(req.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(req.Username, server.config.RefreshTokenDuration)

	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserName:     refreshPayload.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	}
	session, err := server.store.CreateSession(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	res := loginResponse{
		SessionID:             session.ID,
		Token:                 token,
		TokenExpiredAt:        payload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, res)
}

type logoutUserRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (server *Server) logoutUser(ctx *gin.Context) {
	var req logoutUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	err = server.store.DeleteSession(ctx, payload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
}
