package api

import (
	db "drones/db/models"

	"github.com/gin-gonic/gin"
)


type Server struct {
	store  *db.Transaction
	router *gin.Engine
}

func NewServer(store *db.Transaction) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/farm", server.createFarm)
	router.GET("/farm/:farm_code", server.getFarm)
	router.GET("/farm", server.listFarms)

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
