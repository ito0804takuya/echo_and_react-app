package router

import (
	"echo-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()

	// CSRFトークンを使うため、CORSを設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// アクセスを許可するドメイン（フロントエンド）を設定（CSRFトークンはどのフロントを許可するのかを定めて初めて意味があるもの）
		AllowOrigins: []string{
			"http://localhost:3000", 
			os.Getenv("FE_URL"),
		},
		AllowHeaders: []string{
			echo.HeaderOrigin, 
			echo.HeaderContentType, 
			echo.HeaderAccept, 
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods: []string{
			"GET", "PUT", "POST", "DELETE",
		},
		AllowCredentials: true, // cookie送受信のため
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode, // local検証用
		// CookieMaxAge: 60, // 秒
	}))

	// User関連
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

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