package api

import (
	"net/http"
	db "webbanhang/db/sqlc"
	"webbanhang/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) registerHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{"Title": "Register"})
}

type CreateUserRequest struct {
	Name     string `form:"name"`
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateUserRequest
	err := ctx.ShouldBind(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		UserName:     req.Username,
		FullName:     req.Name,
		HashPassword: hashedPassword,
		Email:        req.Email,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
