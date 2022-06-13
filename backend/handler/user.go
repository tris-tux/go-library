package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-library/backend/schema"
	"github.com/tris-tux/go-library/backend/service"
)

type userHandler struct {
	visitorPostgres service.Visitor
}

func NewVisitor(visitorPostgres service.Visitor) *userHandler {
	return &userHandler{visitorPostgres}
}

func (h *userHandler) GetVisitors(c *gin.Context) {
	visitors, err := h.visitorPostgres.FindAll()
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	responseOK(c, http.StatusOK, visitors)
}

func (h *userHandler) GetVisitor(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	visitor, err := h.visitorPostgres.FindByNoIdentitas(id)
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	responseOK(c, http.StatusOK, visitor)
}

func (h *userHandler) CreateVisitor(c *gin.Context) {
	var visitorAddRequest schema.Visitor

	err := c.ShouldBindJSON(&visitorAddRequest)
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	err = h.visitorPostgres.Create(visitorAddRequest)
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	responseOK(c, http.StatusCreated, "success")
}

func (h *userHandler) UpdateVisitor(c *gin.Context) {
	var visitorUpdateRequest schema.Visitor

	err := c.ShouldBindJSON(&visitorUpdateRequest)
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	idString := c.Param("no_identitas")
	id, _ := strconv.ParseUint(idString, 10, 64)

	err = h.visitorPostgres.Update(id, visitorUpdateRequest)
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	responseOK(c, http.StatusOK, "success")
}

func (h *userHandler) DeleteVisitor(c *gin.Context) {
	idString := c.Param("no_identitas")
	id, _ := strconv.ParseUint(idString, 10, 64)

	err := h.visitorPostgres.Delete(id)
	if err != nil {
		responseError(c, ErrorCode(err), err.Error())
		return
	}

	responseOK(c, http.StatusOK, "success")
}

func ErrorCode(er error) int {
	r := er.Error()
	code := r[0:3]
	c, _ := strconv.Atoi(code)
	return c
}

func responseOK(c *gin.Context, code int, body interface{}) {
	c.JSON(code, gin.H{
		"data": body,
	})
}

func responseError(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{
		"message":       "Failed",
		"error_key":     code,
		"error_message": message,
		"error_data":    "{}",
	})
}
