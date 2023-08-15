package api

import (
	"database/sql"
	"drones/db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createFieldRequest struct {
	FieldName      string             `json:"field_name" binding:"required"`
	FieldType      models.FieldTypes  `json:"field_type" binding:"required,oneof= Block GH"`
	FieldFarmID    string             `json:"field_farm_id" binding:"required"`
	FieldVarietyID string             `json:"field_variety_id" binding:"required"`
	FieldPolygon   [][]float32        `json:"field_polygon" binding:"required"`
	FieldArea      string             `json:"field_area" binding:"required"`
	FieldDieback   string             `json:"field_dieback" binding:"required"`
	FieldStageName string             `json:"field_stage_name" binding:"required"`
	FieldStatus    models.FieldStatus `json:"field_status" binding:"required,oneof= active inactive suspended"`
}

func (server *Server) createField(ctx *gin.Context) {
	var req createFieldRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.CreateFieldParams{
		FieldName:      req.FieldName,
		FieldType:      req.FieldType,
		FieldFarmID:    req.FieldFarmID,
		FieldVarietyID: req.FieldVarietyID,
		StGeomfromtext: convertToWKT(req.FieldPolygon),
		FieldArea:      req.FieldArea,
		FieldDieback:   req.FieldDieback,
		FieldStageName: req.FieldStageName,
		FieldStatus:    req.FieldStatus,
	}

	field, err := server.store.CreateField(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, field)

}

type getFieldRequest struct {
	FieldName   string `uri:"field_name" binding:"required"`
	FieldFarmID string `uri:"field_farm_id" binding:"required"`
}

func (server *Server) getField(ctx *gin.Context) {
	var req getFieldRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := models.GetFieldParams{
		FieldName:   req.FieldName,
		FieldFarmID: req.FieldFarmID,
	}
	fmt.Println(req.FieldName)
	field, err := server.store.GetField(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}

	ctx.JSON(http.StatusOK, field)
}

type listFieldRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (server *Server) listField(ctx *gin.Context) {
	var req listFieldRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	field, err := server.store.ListFields(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, field)
}

type updateFieldRequest struct {
	FieldName      string             `json:"field_name" binding:"required"`
	FieldName_2    string             `json:"field_name_to_update" binding:"required"`
	FieldType      models.FieldTypes  `json:"field_type" binding:"required,oneof= Block GH"`
	FieldFarmID    string             `json:"field_farm_id" binding:"required"`
	FieldFarmID_2  string             `json:"field_farm_id_to_update" binding:"required"`
	FieldVarietyID string             `json:"field_variety_id" binding:"required"`
	FieldPolygon   [][]float32        `json:"field_polygon" binding:"required"`
	FieldArea      string             `json:"field_area"`
	FieldDieback   string             `json:"field_dieback" binding:"required"`
	FieldStageName string             `json:"field_stage_name" binding:"required"`
	FieldStatus    models.FieldStatus `json:"field_status" binding:"required,oneof= active inactive suspended"`
}

// type overalFieldUpdateRequest struct {
// 	MyRequest []updateFieldRequest
// }

func (server *Server) updateField(ctx *gin.Context) {
	var req updateFieldRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := models.UpdateFieldParams{
		FieldName:      req.FieldName,
		FieldName_2:    req.FieldName_2,
		FieldType:      req.FieldType,
		FieldFarmID:    req.FieldFarmID,
		FieldFarmID_2:  req.FieldFarmID_2,
		FieldVarietyID: req.FieldVarietyID,
		StGeomfromtext: convertToWKT(req.FieldPolygon),
		FieldArea:      req.FieldArea,
		FieldDieback:   req.FieldDieback,
		FieldStageName: req.FieldStageName,
		FieldStatus:    req.FieldStatus,
	}

	field, err := server.store.UpdateField(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, field)

}
