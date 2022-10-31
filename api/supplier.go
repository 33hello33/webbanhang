package api

import (
	"database/sql"
	"net/http"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createSupplierRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Notes   string `json:"notes"`
}

func (server *Server) createSupplier(ctx *gin.Context) {
	var req createSupplierRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateSupplierParams{
		Name:  req.Name,
		Phone: req.Phone,
		Address: sql.NullString{
			String: req.Address,
			Valid:  len(req.Address) != 0},
		Notes: sql.NullString{
			String: req.Notes,
			Valid:  len(req.Notes) != 0},
	}
	_, err = server.store.CreateSupplier(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type getSupplierRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getSupplierReponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Notes   string `json:"notes"`
}

func (server *Server) getSupplier(ctx *gin.Context) {
	var req getSupplierRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	supplier, err := server.store.GetSupplier(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	res := getSupplierReponse{
		ID:      supplier.ID,
		Name:    supplier.Name,
		Phone:   supplier.Phone,
		Address: supplier.Address.String,
		Notes:   supplier.Notes.String,
	}

	ctx.JSON(http.StatusOK, res)
}

type listSupplierRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (server *Server) supplierHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "supplier.html", gin.H{"supplier": "test"})
}

func (server *Server) listSupplier(ctx *gin.Context) {
	var req listSupplierRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	/*arg := db.ListSupplierParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}*/
	suppliers, err := server.store.ListSupplier(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, suppliers)
}

type deleteSupplierRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteSupplier(ctx *gin.Context) {
	var req deleteSupplierRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = server.store.DeleteSupplier(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type updateSupplierRequest struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Notes   string `json:"notes"`
}

func (server *Server) updateSupplier(ctx *gin.Context) {
	var req updateSupplierRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.UpdateSupplierParams{
		ID:      req.ID,
		Name:    req.Name,
		Address: sql.NullString{String: req.Address, Valid: req.Address != ""},
		Phone:   req.Phone,
		Notes:   sql.NullString{String: req.Notes, Valid: req.Notes != ""},
	}
	_, err = server.store.UpdateSupplier(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
