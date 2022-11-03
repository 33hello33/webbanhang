package api

import (
	"net/http"
	"time"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) invoiceHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "invoice.html", gin.H{"title": "invoice"})
}

func (server *Server) revenueHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "revenue.html", gin.H{"title": "revenue"})
}

type Invoice struct {
	CustomerID int64 `json:"customer_id"`
	TotalMoney int64 `json:"total_money"`
	HadPaid    int64 `json:"had_paid"`
}

type createInvoiceRequest struct {
	Invoice  Invoice         `json:"invoice"`
	Products []db.ProductTbl `json:"products"`
}

func (server *Server) createInvoice(ctx *gin.Context) {
	var req createInvoiceRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	invoiceResult, err := server.store.InvoiceTx(ctx, db.InvoiceTxParams{
		CustomerID: req.Invoice.CustomerID,
		TotalMoney: req.Invoice.TotalMoney,
		HadPaid:    req.Invoice.HadPaid,
		Products:   req.Products,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, invoiceResult)
}

func (server *Server) listInvoice(ctx *gin.Context) {
	invoices, err := server.store.ListInvoice(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, invoices)
}

type findInvoiceRequest struct {
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

func (server *Server) findInvoice(ctx *gin.Context) {
	var req findInvoiceRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	fromDate, err := time.Parse("2006-01-02", req.FromDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	toDate, err := time.Parse("2006-01-02", req.ToDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	invoices, err := server.store.FindInvoice(ctx, db.FindInvoiceParams{
		CreatedAt:   fromDate,
		CreatedAt_2: toDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, invoices)
}
