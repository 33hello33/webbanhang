package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	//remove cache
	for _, product := range req.Products {
		err = server.redisClient.Del(ctx, strconv.FormatInt(product.ID, 10)).Err()
		if err != nil {
			log.Println(err)
		}
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
	FromDate     string `json:"from_date"`
	ToDate       string `json:"to_date"`
	FilterBy     string `json:"filter_by"`
	Filter_Input string `json:"filter_input"`
}

type findInvoiceResponse struct {
	SumTotal int64                       `json:"sum_total"`
	Invoices []db.FindInvoiceFromDateRow `json:"invoices"`
}

func (server *Server) findInvoice(ctx *gin.Context) {
	var req findInvoiceRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	fmt.Println(req)
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

	// add 1 day to find, because date only set by time 00:00:00
	toDate = toDate.AddDate(0, 0, 1)

	invoices, err := server.store.FindInvoiceFromDate(ctx, db.FindInvoiceFromDateParams{
		CreatedAt:   fromDate,
		CreatedAt_2: toDate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	sumTotal, err := server.store.SumToTalFromDate(ctx, db.SumToTalFromDateParams{
		CreatedAt:   fromDate,
		CreatedAt_2: toDate,
	})
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, errResponse(err))
		//return
		sumTotal = 0
		log.Println("errno no row to calc sum")
	}

	res := findInvoiceResponse{
		SumTotal: sumTotal,
		Invoices: invoices,
	}

	ctx.JSON(http.StatusOK, res)
}

type getDetailInvoice struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getDetailInvoice(ctx *gin.Context) {
	var req getDetailInvoice
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	invoiceDetails, err := server.store.GetInvoiceDetail(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, invoiceDetails)
}
