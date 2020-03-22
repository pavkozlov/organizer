package account

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) {
	userRouter := r.Group("/userRouter")
	{
		userRouter.POST("/login", Login)
		userRouter.POST("/reg", Register)
		userRouter.POST("/refresh", RefreshToken)
	}
}
