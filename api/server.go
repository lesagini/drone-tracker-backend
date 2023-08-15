package api

import (
	db "drones/db/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Transaction
	router *gin.Engine
}

func NewServer(store *db.Transaction) *Server {
	server := &Server{store: store}
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "access-control-allow-origin"}
	// router.Use(cors.New(config))
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers,Access-Control-Allow-Origin")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	router.POST("/signup", server.signUpUser)
	router.POST("/login", server.loginUser)

	//FARM
	router.POST("/farm", server.createFarm)
	router.GET("/farm/:farm_code", server.getFarm)
	router.GET("/farm", server.listFarms)

	// FIELD
	router.GET("/field/:field_farm_id/:field_name", server.getField)
	router.GET("/field", server.listField)
	router.POST("/field", server.createField)
	router.PUT("/field", server.updateField)
	router.DELETE("/field")

	server.router = router
	return server
}

// Start runs the HTTP serve on a specific address
func (Server *Server) Start(address string) error {
	return Server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
