package server

import (
	"app/internal/controller"
	mdw "app/internal/middlewares"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Routes(server *gin.Engine, controller *controller.LayerController) {
	admin := server.Group("/users", mdw.Auth2FA(), mdw.AuthRole("admin"))
	{
		admin.POST("", controller.CreateUser)
		admin.GET("", controller.GetUserList)
		admin.GET(":id", controller.GetUser)
		admin.PATCH(":id", controller.EditUser)
		admin.DELETE(":id", controller.DeleteUser)
	}

	auth := server.Group("/auth")
	{
		auth.POST("/login", controller.LoginUser)
		auth.POST("/logout", mdw.Auth2FA(), controller.LogoutUser)
		auth.POST("/register", controller.RegisterUser)
		auth.POST("/refresh-token")
		auth.POST("/verify-email")
		forgetPassword := auth.Group("/forget-password")
		{
			forgetPassword.GET("/send-token/:email", controller.SendRefreshForEmail)
			forgetPassword.POST("/refresh-password", controller.RefreshPassword)
		}
	}

	me := server.Group("/me", mdw.Auth2FA())
	{
		me.GET("", controller.GetMyInfo)
		me.PATCH("", controller.EditMyInfo)
	}
}

func HealthCheck(server *gin.Engine, dbSQL *sql.DB, Redis *redis.Client) {
	server.GET("/health", func(ctx *gin.Context) {
		errList := map[string]any{
			"db-sql": true,
			"redis":  true,
		}
		if err := dbSQL.Ping(); err != nil {
			errList["db-sql"] = false
		}
		if err := Redis.Ping(context.Background()).Err(); err != nil {
			errList["redis"] = false
		}
		status := 200
		for _, y := range errList {
			if y == false {
				status = 500
			}
		}

		ctx.JSON(status, errList)
	})
}
