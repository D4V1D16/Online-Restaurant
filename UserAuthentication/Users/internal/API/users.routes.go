package API

import (
	controllers "userAuth/Users/internal/Controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	api := router.Group("/users")

	api.GET("/", controllers.GetAllUser)

	api.POST("/", controllers.CreateUser)

	api.DELETE("/:id", controllers.DeleteUser)

	api.GET("/:id", controllers.GetSingleUser)
	/*


		api.PATCH("/:id", updateUser)


	*/

}
