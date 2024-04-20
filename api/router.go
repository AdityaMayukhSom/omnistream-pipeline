package api

import (
	"fmt"

	authControllers "devstream.in/pixelated-pipeline/api/controllers/auth"
	postControllers "devstream.in/pixelated-pipeline/api/controllers/posts"
	templatesControllers "devstream.in/pixelated-pipeline/api/controllers/templates"
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
	er.echo.GET("/hello-world", templatesControllers.HelloWorld)

	apiV1 := er.echo.Group("/api/v1")

	apiV1.GET("/posts", postControllers.ReturnAllPosts)
	apiV1.GET("/posts/:username/:postname", postControllers.ReturnSinglePost)

	authRoute := apiV1.Group("/auth")
	authRoute.POST("/register", authControllers.SignUp)
	authRoute.POST("/login", authControllers.LogIn)
	authRoute.POST("/refresh", authControllers.Refresh)

	restrictedRoute := apiV1.Group("/")
	restrictedRoute.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	restrictedRoute.POST("/posts", postControllers.CreatePost)
	restrictedRoute.DELETE("/posts/{id}", postControllers.DeletePost)
	restrictedRoute.PUT("/posts/{id}", postControllers.UpdatePost)
}

func (er *EchoRouter) Start() {
	address := fmt.Sprintf(":%d", config.GetPort())
	er.echo.Logger.Fatal(er.echo.Start(address))
}
