package http

import (
	"errors"
	"strconv"

	"github.com/Crunchy89/go-mysql/app/user/service"
	"github.com/Crunchy89/go-mysql/utils/s"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	User service.UserService
}

func (h *UserHandler) GetAll(c *gin.Context) {
	res, err := h.User.GetAll()
	s.Auto(c, res, err)
}
func (h *UserHandler) GetUserById(c *gin.Context) {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		c.AbortWithError(422, errors.New("id not valid"))
	}
	res, err := h.User.GetById(id)
	s.Auto(c, res, err)
}
