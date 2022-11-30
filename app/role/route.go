package role

import (
	_roleHttp "github.com/Crunchy89/go-mysql/app/role/http"
	_roleRepo "github.com/Crunchy89/go-mysql/app/role/repository"
	_roleService "github.com/Crunchy89/go-mysql/app/role/service"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func RoleRoute(handler *gin.Engine, database *gorm.DB) {
	roleRepository := _roleRepo.NewRoleRepository(database)
	roleService := _roleService.NewRoleService(roleRepository)
	roleHttp := _roleHttp.RoleHandler{
		Role: roleService,
	}
	v1 := handler.Group("api").Group("v1")
	role := v1.Group("role")
	role.GET("", roleHttp.GetAll)
	role.GET("/:id", roleHttp.GetRoleById)
	role.POST("", roleHttp.CreateRole)
	role.PUT("", roleHttp.UpdateRole)
	role.DELETE("", roleHttp.DeleteRole)
}
