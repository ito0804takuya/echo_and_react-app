package router

import (
	"echo-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()

	// User関連
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	// Task関連
	// routerグループを作る
	t := e.Group("/tasks")
	// taskのrouterグループに対してjwtトークン用のミドルウェアを用意
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")), // 暗号化したキー
		TokenLookup: "cookie:token", // tokenという名前のcookie
	}))
	// Group("/tasks")でグループ化されているので、GET("") は /tasks のこと
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteeTask)

	return e
}