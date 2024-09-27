package API


import(
	"github.com/gin-gonic/gin"

	"userAuth/Users/internal/Controllers"
)


func ProfileRoutes(router *gin.Engine) {
	api := router.Group("/profiles")

	api.GET("/", controllers.GetAllProfiles)

	api.POST("/", controllers.CreateProfile)

	api.DELETE("/:id", controllers.DeleteProfile)

	api.GET("/:id", controllers.GetSingleProfile)
	
	api.PATCH("/:id", controllers.UpdateProfile)

}
