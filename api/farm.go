package api

import (
	"database/sql"
	"drones/db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type returnJsonRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) returnJson(ctx *gin.Context) {
	var req returnJsonRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req.Email)

	ctx.JSON(http.StatusOK, req)

}

type createFarmRequest struct {
	FarmCode        string `json:"farm_code" binding:"required"`
	FarmCoordinates string `json:"farm_coordinates" binding:"required"`
	FarmAirspace    string `json:"farm_airspace" binding:"required"`
	FarmLocation    string `json:"farm_location" binding:"required,oneof= Naivasha Nanyuki"`
	FarmGeolocation string `json:"farm_geolocation" binding:"required"`
	FarmContact     int64  `json:"farm_contact" binding:"required"`
}

func (server *Server) createFarm(ctx *gin.Context) {
	var req createFarmRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.CreateFarmParams{
		FarmCode:     req.FarmCode,
		FarmAirspace: req.FarmAirspace,
		FarmLocation: req.FarmLocation,
		FarmContact:  req.FarmContact,
	}

	farm, err := server.store.CreateFarm(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, farm)

}

type getFarmRequest struct {
	FarmCode string `uri:"farm_code" binding:"required"`
}

func (server *Server) getFarm(ctx *gin.Context) {
	var req getFarmRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println(req.FarmCode)
	farm, err := server.store.GetFarm(ctx, req.FarmCode)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}

	ctx.JSON(http.StatusOK, farm)
}

type listFarmsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFarms(ctx *gin.Context) {
	var req listFarmsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := models.ListFarmsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	farm, err := server.store.ListFarms(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, farm)
}
