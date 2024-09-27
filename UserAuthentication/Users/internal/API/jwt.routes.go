package API


import(
	"github.com/gin-gonic/gin"

	"userAuth/Users/internal/Controllers"
)


func JWTRoutes(router *gin.Engine) {
	api := router.Group("/jwt")

	api.POST("/login",controllers.Login)
}