package api

import (
	"fmt"

	"devstream.in/pixelated-pipeline/api/controllers"
	"devstream.in/pixelated-pipeline/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

type Router interface {
	RegisterRoutes()
	Start()
}

func NewRouter() Router {
	return NewEchoRouter()
}

// ------------------------------------------------------------------------

type EchoRouter struct {
	echo *echo.Echo
}

func NewEchoRouter() *EchoRouter {
	return &EchoRouter{
		echo: echo.New(),
	}
}

func (er *EchoRouter) RegisterRoutes() {
	er.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderContentLength,
		},
	}))

	er.echo.Static("/", "static")

	er.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	er.echo.Renderer = NewRenderer("./views/*", true)
	er.echo.GET("/hello-world", controllers.HelloWorld)

	apiV1 := er.echo.Group("/api/v1")

	apiV1.GET("/posts", controllers.ReturnAllPosts)
	apiV1.GET("/posts/:username/:postname", controllers.ReturnSinglePost)

	authRoute := apiV1.Group("/auth")
	authRoute.POST("/register", controllers.SignUp)
	authRoute.POST("/login", controllers.LogIn)
	authRoute.POST("/refresh", controllers.Refresh)
	authRoute.GET("/logout", controllers.LogOut)

	restrictedRoute := apiV1.Group("/")
	restrictedRoute.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	restrictedRoute.POST("/posts", controllers.CreatePost)
	restrictedRoute.DELETE("/posts/{id}", controllers.DeletePost)
	restrictedRoute.PUT("/posts/{id}", controllers.UpdatePost)
}

func (er *EchoRouter) Start() {
	address := fmt.Sprintf(":%d", config.GetPort())
	er.echo.Logger.Fatal(er.echo.Start(address))
}
