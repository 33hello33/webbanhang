package api

import (
	"database/sql"
	"net/http"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	PriceImport int64  `json:"price_import"`
	Amount      int64  `json:"amount"`
	Price       int64  `json:"price"`
	WareHouse   string `json:"warehouse"`
	IdSupplier  int64  `json:"id_supplier"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateProductParams{
		Name:        req.Name,
		Unit:        req.Unit,
		Price:       req.Price,
		PriceImport: req.PriceImport,
		Amount:      req.Amount,
		Warehouse:   req.WareHouse,
		IDSupplier:  sql.NullInt64{Int64: req.IdSupplier, Valid: req.IdSupplier != 0},
	}

	product, err := server.store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type listProductRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (server *Server) productHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "product.html", gin.H{"title": "product"})
}

func (server *Server) listProduct(ctx *gin.Context) {
	var req listProductRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	//arg := db.ListProductsParams{
	//	Limit:  int32(req.Limit),
	//	Offset: int32(req.Offset),
	//}
	products, err := server.store.ListProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

type deleteProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = server.store.DeleteProduct(ctx, int64(req.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getProductResponse struct {
	Name        string `json:"name"`
	ID          int64  `json:"id"`
	Unit        string `json:"unit"`
	PriceImport int64  `json:"price_import"`
	Amount      int64  `json:"amount"`
	Price       int64  `json:"price"`
	WareHouse   string `json:"warehouse"`
	IdSupplier  int64  `json:"id_supplier"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	res := getProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Unit:        product.Unit,
		PriceImport: product.PriceImport,
		Amount:      product.Amount,
		Price:       product.Price,
		WareHouse:   product.Warehouse,
		IdSupplier:  product.IDSupplier.Int64,
	}
	ctx.JSON(http.StatusOK, res)
}

type updateProductRequest struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	PriceImport int64  `json:"price_import"`
	Amount      int64  `json:"amount"`
	Price       int64  `json:"price"`
	WareHouse   string `json:"warehouse"`
	IdSupplier  int64  `json:"id_supplier"`
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var req updateProductRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err = server.store.UpdateProduct(ctx, db.UpdateProductParams{
		ID:          req.ID,
		Amount:      req.Amount,
		Price:       req.Amount,
		PriceImport: req.PriceImport,
		Warehouse:   req.WareHouse,
		IDSupplier:  sql.NullInt64{Int64: req.IdSupplier, Valid: req.IdSupplier != 0},
		Unit:        req.Unit,
		Name:        req.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type searchProductRequest struct {
	Name string `uri:"name"`
}

func (server *Server) searchProduct(ctx *gin.Context) {
	var req searchProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	products, err := server.store.SearchProductLikeName(ctx, "%"+req.Name+"%")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}
