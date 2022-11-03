package api

import (
	"net/http"
	db "webbanhang/db/sqlc"
	"webbanhang/token"
	"webbanhang/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store      *db.Store
	route      *gin.Engine
	tokenMaker *token.PasetoMaker
	config     util.Config
}

func NewServer(store *db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.SetupRoute()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func (server *Server) SetupRoute() {
	router := gin.Default()
	router.LoadHTMLGlob("./static/html/*")
	router.Static("/assets/", "./static/assets")

	router.GET("/test", server.test)

	router.GET("/", server.invoiceHandler)

	router.GET("/login", server.loginHandler)
	router.GET("/register", server.registerHandler)
	router.POST("/login", server.loginUser)
	router.POST("/logout", server.logoutUser)
	router.POST("/register", server.createAccount)
	router.GET("/token/renew", server.renewTokenHandler)
	router.POST("/token/renew", server.renewToken)

	authRoutes := router.Group("/").Use(authMiddleware(*server.tokenMaker))

	authRoutes.GET("/product", server.productHandler)
	authRoutes.GET("/product/:id", server.getProduct)
	authRoutes.POST("/product/create", server.createProduct)
	authRoutes.GET("/product/list", server.listProduct)
	authRoutes.PUT("/product/:id", server.updateProduct)
	authRoutes.DELETE("/product/:id", server.deleteProduct)

	authRoutes.GET("/supplier", server.supplierHandler)
	authRoutes.POST("/supplier/create", server.createSupplier)
	authRoutes.GET("/supplier/list", server.listSupplier)
	authRoutes.GET("/supplier/:id", server.getSupplier)
	authRoutes.PUT("/supplier/:id", server.updateSupplier)
	authRoutes.DELETE("/supplier/:id", server.deleteSupplier)

	authRoutes.GET("/invoice", server.invoiceHandler)
	authRoutes.POST("/invoice/create", server.createInvoice)
	authRoutes.GET("/invoice/list", server.listInvoice)
	authRoutes.POST("/invoice/find", server.findInvoice)

	authRoutes.GET("/revenue", server.revenueHandler)

	authRoutes.GET("/customer", server.customerHandler)
	authRoutes.POST("/customer/create", server.createCustomer)
	authRoutes.GET("/customer/list", server.listCustomer)
	authRoutes.GET("/customer/:id", server.getDetailCustomer)
	authRoutes.DELETE("/customer/:id", server.deleteCustomer)
	authRoutes.PUT("/customer/:id", server.updateCustomer)

	server.route = router
}

func errResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}

func (server *Server) test(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "test.html", nil)
}
