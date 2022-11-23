package api

import (
	db "webbanhang/db/sqlc"
	"webbanhang/token"
	"webbanhang/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	store       *db.Store
	route       *gin.Engine
	tokenMaker  *token.PasetoMaker
	config      util.Config
	redisClient *redis.Client
}

func NewServer(store *db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, err
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
	})
	server := &Server{
		store:       store,
		config:      config,
		tokenMaker:  tokenMaker,
		redisClient: redisClient,
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
	authRoutes.GET("/product/:id", server.getCacheProduct(), server.getProduct) // add test middleware
	authRoutes.POST("/product/create", server.createProduct)
	authRoutes.GET("/product/list", server.listProduct)
	authRoutes.PUT("/product/:id", server.updateProduct)
	authRoutes.DELETE("/product/:id", server.deleteProduct)
	authRoutes.GET("/product/search/:name", server.searchProduct)
	authRoutes.POST("/product/copy/:id", server.copyProduct)
	authRoutes.POST("/product/import_from_file", server.importProductFromFile)
	authRoutes.GET("/product/export_to_file", server.exportProductToFile)

	authRoutes.GET("/supplier", server.supplierHandler)
	authRoutes.POST("/supplier/create", server.createSupplier)
	authRoutes.GET("/supplier/list", server.listSupplier)
	authRoutes.GET("/supplier/:id", server.getSupplier)
	authRoutes.PUT("/supplier/:id", server.updateSupplier)
	authRoutes.DELETE("/supplier/:id", server.deleteSupplier)
	authRoutes.GET("/supplier/search/:name", server.searchSupplier)

	authRoutes.GET("/invoice", server.invoiceHandler)
	authRoutes.POST("/invoice/create", server.createInvoice)
	authRoutes.GET("/invoice/list", server.listInvoice)
	authRoutes.POST("/invoice/find", server.findInvoice)
	authRoutes.GET("/invoice/:id", server.getInvoice)
	authRoutes.GET("/invoice/detail/:id", server.getDetailInvoice)
	authRoutes.POST("/invoice/update/:id", server.updateInvoice)
	authRoutes.GET("/invoice/print/:id", server.printInvoiceHandler)

	authRoutes.GET("/revenue", server.revenueHandler)

	authRoutes.GET("/customer", server.customerHandler)
	authRoutes.POST("/customer/create", server.createCustomer)
	authRoutes.GET("/customer/list", server.listCustomer)
	authRoutes.GET("/customer/:id", server.getDetailCustomer)
	authRoutes.DELETE("/customer/:id", server.deleteCustomer)
	authRoutes.PUT("/customer/:id", server.updateCustomer)
	authRoutes.GET("/customer/search/:name", server.searchCustomer)

	authRoutes.GET("/statistic", server.statisticHandler)
	authRoutes.POST("/statistic/top_selling_product", server.getTopSellingProduct)
	server.route = router
}

func errResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}
