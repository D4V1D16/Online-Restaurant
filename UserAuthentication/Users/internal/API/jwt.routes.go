package API


import(
	"github.com/gin-gonic/gin"

	"userAuth/Users/internal/Controllers"
)


func JWTRoutes(router *gin.Engine) {
	api := router.Group("/jwt")

	api.POST("/login",controllers.Login)
	api.GET("/protected",controllers.ProtectedRoute)
	api.GET("/refresh",controllers.Refresh)
	api.GET("/logout",controllers.Logout)
}