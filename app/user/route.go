package user

import (
	_userHttp "github.com/Crunchy89/go-mysql/app/user/http"
	_userRepo "github.com/Crunchy89/go-mysql/app/user/repository"
	_userService "github.com/Crunchy89/go-mysql/app/user/service"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func UserRoute(handler *gin.Engine, database *gorm.DB) {
	userRepository := _userRepo.NewUserRepository(database)
	userService := _userService.NewUserService(userRepository)
	userHttp := _userHttp.UserHandler{
		User: userService,
	}
	v1 := handler.Group("api").Group("v1")
	user := v1.Group("user")
	user.GET("", userHttp.GetAll)
	user.GET("/:id", userHttp.GetUserById)
	user.POST("", userHttp.CreateUser)
	user.PUT("", userHttp.UpdateUser)
	user.DELETE("", userHttp.DeleteUser)
}
