package api

import (
	"database/sql"
	"net/http"
	"strings"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) customerHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "customer.html", gin.H{"title": "customer"})
}

type createCustomer struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (server *Server) createCustomer(ctx *gin.Context) {
	var req createCustomer
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if req.Phone == "" {
		req.Phone = "0"
	}
	if req.Name == "" {
		req.Name = "Khách vãng lai"
	}

	arg := db.CreateCustomerParams{
		Name:    req.Name,
		Phone:   req.Phone,
		Address: sql.NullString{String: req.Address, Valid: req.Address != ""},
	}
	_, err = server.store.CreateCustomer(ctx, arg)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") == true {
			ctx.JSON(http.StatusOK, nil)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) listCustomer(ctx *gin.Context) {
	customers, err := server.store.ListCustomer(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

type getDetailCustomerRequest struct {
	Phone string `uri:"phone"`
}

func (server *Server) getDetailCustomer(ctx *gin.Context) {
	var req getDetailCustomerRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	customers, err := server.store.GetCustomer(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}
