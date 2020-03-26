package account

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", Login)
		userRouter.POST("/reg", Register)
		userRouter.POST("/refresh", RefreshToken)
	}

	profileRouter := r.Group("/user")
	profileRouter.Use(JWTAuth())
	{
		profileRouter.GET("/me", Profile)
	}
}
