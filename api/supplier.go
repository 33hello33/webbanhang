package api

import (
	"database/sql"
	"net/http"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) supplierHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "supplier.html", gin.H{"supplier": "test"})
}

type createSupplierRequest struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Zalo       string `json:"zalo"`
	Notes      string `json:"notes"`
	BankName   string `json:"bank_name"`
	BankNumber string `json:"bank_number"`
}

func (server *Server) createSupplier(ctx *gin.Context) {
	var req createSupplierRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	if req.Phone == "" {
		req.Phone = "0"
	}
	if req.Name == "" {
		req.Name = "Nhập lẻ"
	}

	supplier, err := server.store.GetSupplierByPhone(ctx, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			arg := db.CreateSupplierParams{
				Name:  req.Name,
				Phone: req.Phone,
				Address: sql.NullString{
					String: req.Address,
					Valid:  len(req.Address) != 0},
				Notes: sql.NullString{
					String: req.Notes,
					Valid:  len(req.Notes) != 0},
				Zalo: sql.NullString{
					String: req.Zalo,
					Valid:  len(req.Zalo) != 0},
				BankName: sql.NullString{
					String: req.BankName,
					Valid:  len(req.BankName) != 0},
				BankNumber: sql.NullString{
					String: req.BankNumber,
					Valid:  len(req.BankNumber) != 0},
			}

			supplier, err = server.store.CreateSupplier(ctx, arg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return
			}

		} else {
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, supplier)
}

type getSupplierRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getSupplierReponse struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Zalo       string `json:"zalo"`
	Notes      string `json:"notes"`
	BankName   string `json:"bank_name"`
	BankNumber string `json:"bank_number"`
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
		ID:         supplier.ID,
		Name:       supplier.Name,
		Phone:      supplier.Phone,
		Address:    supplier.Address.String,
		Notes:      supplier.Notes.String,
		Zalo:       supplier.Zalo.String,
		BankName:   supplier.BankName.String,
		BankNumber: supplier.BankNumber.String,
	}

	ctx.JSON(http.StatusOK, res)
}

type listSupplierRequest struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
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
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Notes      string `json:"notes"`
	Zalo       string `json:"zalo"`
	BankNumber string `json:"bank_number"`
	BankName   string `json:"bank_name"`
}

func (server *Server) updateSupplier(ctx *gin.Context) {
	var req updateSupplierRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.UpdateSupplierParams{
		ID:         req.ID,
		Name:       req.Name,
		Address:    sql.NullString{String: req.Address, Valid: req.Address != ""},
		Phone:      req.Phone,
		Notes:      sql.NullString{String: req.Notes, Valid: req.Notes != ""},
		Zalo:       sql.NullString{String: req.Zalo, Valid: req.Zalo != ""},
		BankName:   sql.NullString{String: req.BankName, Valid: req.BankName != ""},
		BankNumber: sql.NullString{String: req.BankNumber, Valid: req.BankNumber != ""},
	}
	_, err = server.store.UpdateSupplier(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type searchSupplierRequest struct {
	Name string `uri:"name"`
}

func (server *Server) searchSupplier(ctx *gin.Context) {
	var req searchSupplierRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	suppliers, err := server.store.SearchSupplierLikeName(ctx, "%"+req.Name+"%")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, suppliers)
}
