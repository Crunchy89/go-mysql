package http

import (
	"errors"
	"strconv"

	"github.com/Crunchy89/go-mysql/app/role/payload"
	"github.com/Crunchy89/go-mysql/app/role/service"
	"github.com/Crunchy89/go-mysql/utils/s"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	Role service.RoleService
}

func (h *RoleHandler) GetAll(c *gin.Context) {
	res, err := h.Role.GetAll()
	s.Auto(c, res, err)
}
func (h *RoleHandler) GetRoleById(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.AbortWithError(422, errors.New("id not valid"))
	}
	res, err := h.Role.GetById(id)
	s.Auto(c, res, err)
}
func (h *RoleHandler) CreateRole(c *gin.Context) {
	role := new(payload.RoleCreate)
	if err := c.BindJSON(&role); err != nil {
		s.AbortWithStatus(c, 422, err)
		return
	}
	res, err := h.Role.CreateRole(role)
	s.Auto(c, res, err)
}
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	role := new(payload.RoleUpdate)
	if err := c.BindJSON(&role); err != nil {
		s.AbortWithStatus(c, 422, err)
		return
	}
	res, err := h.Role.UpdateRole(role)
	s.Auto(c, res, err)
}
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	role := new(payload.RoleDelete)
	if err := c.BindJSON(&role); err != nil {
		s.AbortWithStatus(c, 422, err)
		return
	}
	res, err := h.Role.DeleteRole(role)
	s.Auto(c, res, err)
}
