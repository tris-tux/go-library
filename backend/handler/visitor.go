package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-library/backend/helper"
	"github.com/tris-tux/go-library/backend/schema"
	"github.com/tris-tux/go-library/backend/service"
)

type VisitorHandler interface {
	Update(c *gin.Context)
	Profile(c *gin.Context)
}

type visitorHandler struct {
	visitorService service.VisitorService
	jwtService     service.JWTService
}

//VisitorHandler is creating anew instance of VisitorControlller
func NewVisitorHandler(visitorService service.VisitorService, jwtService service.JWTService) VisitorHandler {
	return &visitorHandler{
		visitorService: visitorService,
		jwtService:     jwtService,
	}
}

func (h *visitorHandler) Update(c *gin.Context) {
	var visitorUpdateDTO schema.VisitorUpdateDTO
	errDTO := c.ShouldBindJSON(&visitorUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := c.GetHeader("Authorization")
	authorizationHeader := c.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		helper.BuildErrorResponse("Invalid token", "Bareer", http.StatusBadRequest)
		return
	}

	authHeader := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, errToken := h.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["visitor_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	visitorUpdateDTO.ID = id
	u := h.visitorService.Update(visitorUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	c.JSON(http.StatusOK, res)
}

func (h *visitorHandler) Profile(c *gin.Context) {
	// authHeader := c.GetHeader("Authorization")
	authorizationHeader := c.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		helper.BuildErrorResponse("Invalid token", "Bareer", http.StatusBadRequest)
		return
	}

	authHeader := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["visitor_id"])
	visitor := h.visitorService.Profile(id)
	res := helper.BuildResponse(true, "OK", visitor)
	c.JSON(http.StatusOK, res)

}
