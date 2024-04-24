package api

import (
	"log"
	"net/http"
	"time"

	"github.com/a-pavithraa/product-service-golang/models" // Import the models package
	"github.com/a-pavithraa/product-service-golang/service"
	"github.com/a-pavithraa/product-service-golang/util"
	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-glog"
	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
)

type Server struct {
	Router *gin.Engine
}

func NewApiServer(settings util.AppSettings) *Server {
	router := gin.Default()
	server := &Server{
		Router: router,
	}
	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(ginkeycloak.RequestLogger([]string{"uid"}, "data"))
	router.Use(gin.Recovery())
	var sbbEndpoint = ginkeycloak.KeycloakConfig{
		Url:           "http://keycloak:8080",
		Realm:         "Products",
		FullCertsPath: nil,
	}
	privateGroup := router.Group("/api")
	privateGroup.Use(ginkeycloak.Auth(ginkeycloak.AuthCheck(), sbbEndpoint))
	privateGroup.GET("/products", server.GetProducts)
	privateGroup.GET("/product/:id", server.GetProductByID)
	privateGroup.POST("/product", server.PostProduct)

	return server
}

func (server *Server) Start(addr string) error {
	return server.Router.Run(addr)
}
func (server *Server) GetProducts(c *gin.Context) {
	products, err := service.GetProducts()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
	}

	// Return the products as JSON response
	c.JSON(http.StatusOK, products)
}
func (server *Server) GetProductByID(c *gin.Context) {
	// Get the product ID from the request parameters
	productID := c.Param("id")
	log.Println("Product ID:", productID)

	product, err := service.GetProductByID(productID)
	if _, ok := err.(*service.ProductNotFoundError); ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	// Return the product as JSON response
	c.JSON(http.StatusOK, product)
}
func (server *Server) PostProduct(c *gin.Context) {
	// Parse the request body into a Product struct
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := service.CheckProductNameExists(product.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name already exists"})
		return
	}
	err = service.PostProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product"})
		return
	}
	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}
