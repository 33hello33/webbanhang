package api

import (
	"database/sql"
	"net/http"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) customerHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "customer.html", gin.H{"title": "customer"})
}

type StringNull struct {
	String string `json:"String"`
	Valid  bool   `json:"Valid"`
}

type createCustomerRequest struct {
	Name    string     `json:"name"`
	Phone   string     `json:"phone"`
	Address StringNull `json:"address"`
}

func (server *Server) createCustomer(ctx *gin.Context) {
	var customer db.Customer
	var req createCustomerRequest
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

	customer, err = server.store.GetCustomerByPhone(ctx, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			arg := db.CreateCustomerParams{
				Name:    req.Name,
				Phone:   req.Phone,
				Address: sql.NullString(req.Address),
			}

			customer, err = server.store.CreateCustomer(ctx, arg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, customer)
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
	ID int64 `uri:"id"`
}

func (server *Server) getDetailCustomer(ctx *gin.Context) {
	var req getDetailCustomerRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	customers, err := server.store.GetCustomer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

type deleteCustomerRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteCustomer(ctx *gin.Context) {
	var req deleteCustomerRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = server.store.DeleteCustomer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type updateCustomerRequest struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Phone   string     `json:"phone"`
	Address StringNull `json:"address"`
}

func (server *Server) updateCustomer(ctx *gin.Context) {
	var req updateCustomerRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err = server.store.UpdateCustomer(ctx, db.UpdateCustomerParams{
		ID:      req.ID,
		Name:    req.Name,
		Address: sql.NullString(req.Address),
		Phone:   req.Phone,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type searchCustomerRequest struct {
	Name string `uri:"name"`
}

func (server *Server) searchCustomer(ctx *gin.Context) {
	var req searchCustomerRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	customers, err := server.store.SearchCustomerLikeName(ctx, "%"+req.Name+"%")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}
