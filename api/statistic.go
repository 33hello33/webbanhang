package api

import (
	"fmt"
	"net/http"
	"time"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) statisticHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "statistic.html", nil)
}

type getTopSellingProductRequest struct {
	NumberTop int64  `json:"number_top"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
}

func (server *Server) getTopSellingProduct(ctx *gin.Context) {
	var req getTopSellingProductRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	fromDate, err := time.Parse("2006-01-02", req.FromDate)
	if err != nil {
		err = fmt.Errorf("%w, cannot parse date time: fromdate", err)
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	toDate, err := time.Parse("2006-01-02", req.ToDate)
	if err != nil {
		err = fmt.Errorf("%w, cannot parse date time: ToDate", err)
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// add 1 day to find, because date only set by time 00:00:00
	toDate = toDate.AddDate(0, 0, 1)
	top, err := server.store.GetTopSellingProducts(ctx, db.GetTopSellingProductsParams{
		CreatedFrom: fromDate,
		CreatedTo:   toDate,
		NumberTop:   int32(req.NumberTop),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, top)
}
