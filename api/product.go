package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
	db "webbanhang/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (server *Server) productHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "product.html", gin.H{"title": "product"})
}

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
		IDSupplier:  req.IdSupplier,
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
		IdSupplier:  product.IDSupplier,
	}

	// convert product to byte and set to redis
	pd, err := json.Marshal(product)
	if err != nil {
		log.Println("cant marshaling the data product")
	}

	err = server.redisClient.Set(ctx, strconv.FormatInt(product.ID, 10), pd, 10*60*time.Second).Err()
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, res)
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var req getProductResponse
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err = server.store.UpdateProduct(ctx, db.UpdateProductParams{
		ID:          req.ID,
		Amount:      req.Amount,
		Price:       req.Price,
		PriceImport: req.PriceImport,
		Warehouse:   req.WareHouse,
		IDSupplier:  req.IdSupplier,
		Unit:        req.Unit,
		Name:        req.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	//remove cache
	err = server.redisClient.Del(ctx, strconv.FormatInt(req.ID, 10)).Err()
	if err != nil {
		log.Println(err)
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

func (server *Server) getCacheProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req getProductRequest
		err := ctx.ShouldBindUri(&req)
		if err != nil {
			ctx.Next()
		}

		bytes, err := server.redisClient.Get(ctx, strconv.FormatInt(req.ID, 10)).Bytes()

		var product getProductResponse

		err = json.Unmarshal(bytes, &product)
		if err != nil {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusOK, product)
			return
		}
		ctx.Next()
	}
}

type copyProductRequest struct {
	ID int64 `uri:"id"`
}

func (server *Server) copyProduct(ctx *gin.Context) {
	var req copyProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	product, err := server.store.CopyProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (server *Server) importProductFromFile(ctx *gin.Context) {
	formFile, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	openFile, err := formFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	csvLines, err := csv.NewReader(openFile).ReadAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	for _, line := range csvLines {
		PriceImport, err := strconv.ParseInt(line[3], 10, 64)
		if err != nil {
			continue
		}

		Amount, err := strconv.ParseInt(line[4], 10, 64)
		if err != nil {
			continue
		}

		Price, err := strconv.ParseInt(line[5], 10, 64)
		if err != nil {
			continue
		}

		IdSupplier, err := strconv.ParseInt(line[8], 10, 64)
		if err != nil {
			continue
		}

		arg := db.CreateProductParams{
			Name:        line[1],
			Unit:        line[2],
			PriceImport: PriceImport,
			Price:       Price,
			Amount:      Amount,
			Warehouse:   line[6],
			IDSupplier:  IdSupplier,
		}
		_, err = server.store.CreateProduct(ctx, arg)
		if err != nil {
			log.Println(line)
			log.Println(err)
		}
	}
	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) exportProductToFile(ctx *gin.Context) {

	cmd := "docker exec postgres12 psql -d webbanhang -c \"COPY products to stdout csv header\""

	//  export file products.csv to docker
	err, out, _ := execCommandline(cmd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Data(http.StatusOK, "application/vnd.ms-excel", []byte(out.String()))
}

func execCommandline(command string) (err error, out bytes.Buffer, stderr bytes.Buffer) {
	fmt.Println(command)
	cmd := exec.Command("powershell", command)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	fmt.Println("Result: " + out.String())
	return
}
