package api

import (
	"database/sql"
	"drones/db/models"
	"drones/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "sagini"

type signUpUserRequest struct {
	Email    string            `json:"email" binding:"required"`
	Username string            `json:"username" binding:"required"`
	Password string            `json:"password" binding:"required"`
	Access   models.UserAccess `json:"access" binding:"required,oneof= admin user"`
}

func (server *Server) signUpUser(ctx *gin.Context) {
	var req signUpUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashed_password, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	arg := models.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashed_password,
		Access:   req.Access,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user.Username)

}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Username   string
	Logintoken string
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginRequest
	var res loginResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, "invalid credentials")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.SetCookie("jwt", token, int(time.Now().Add(time.Hour*24).Second()), "/", "localhost", false, true)
	res.Logintoken = token
	res.Username = req.Username
	ctx.JSON(http.StatusOK, res)
}
