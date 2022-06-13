package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-library/backend/helper"
	"github.com/tris-tux/go-library/backend/schema"
	"github.com/tris-tux/go-library/backend/service"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService service.AuthService, jwtService service.JWTService) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (h *authHandler) Login(c *gin.Context) {
	var loginDTO schema.LoginDTO
	errDTO := c.ShouldBindJSON(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := h.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(schema.Visitor); ok {
		generatedToken := h.jwtService.GenerateToken(strconv.FormatUint(v.NoIdentitas, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (h *authHandler) Register(c *gin.Context) {
	var registerDTO schema.RegisterDTO
	errDTO := c.ShouldBindJSON(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !h.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		c.JSON(http.StatusConflict, response)
	} else {
		createdVisitor := h.authService.CreateVisitor(registerDTO)
		token := h.jwtService.GenerateToken(strconv.FormatUint(createdVisitor.NoIdentitas, 10))
		createdVisitor.Token = token
		response := helper.BuildResponse(true, "OK!", createdVisitor)
		c.JSON(http.StatusCreated, response)
	}
}
