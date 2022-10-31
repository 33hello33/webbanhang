package api

import (
	"net/http"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) invoiceHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "invoice.html", gin.H{"title": "invoice"})
}

func (server *Server) revenueHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "revenue.html", gin.H{"title": "revenue"})
}

type createInvoiceRequest struct {
	Invoice  Invoice         `json:"invoice"`
	Products []db.ProductTbl `json:"products"`
}

type Invoice struct {
	CustomerPhone string `json:"customer_phone"`
	TotalMoney    int64  `json:"total_money"`
	HadPaid       int64  `json:"had_paid,string"`
}

func (server *Server) createInvoice(ctx *gin.Context) {
	var req createInvoiceRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	invoiceResult, err := server.store.InvoiceTx(ctx, db.InvoiceTxParams{
		CustomersPhone: req.Invoice.CustomerPhone,
		TotalMoney:     req.Invoice.TotalMoney,
		HadPaid:        req.Invoice.HadPaid,
		Products:       req.Products,
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
