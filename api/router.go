package api

import (
	"fmt"
	"os"

	"devstream.in/pixelated-pipeline/api/controllers"
	"devstream.in/pixelated-pipeline/api/middlewares"
	"devstream.in/pixelated-pipeline/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// echoSwagger "github.com/swaggo/echo-swagger"
	// _ "github.com/swaggo/echo-swagger/example/docs"
)

type Router interface {
	RegisterRoutes()
	Start()
}

func NewRouter() Router {
	return NewEchoRouter()
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

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

	// To declare global router wise static directory.
	// er.echo.Static("/", "static")

	// For documentation generation for API, currently not required
	// er.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	er.registerWebRoutes()
	er.registerApiRoutes()
}

func (er *EchoRouter) registerWebRoutes() {
	er.echo.Renderer = controllers.NewRenderer("./views/*", true)

	// Group made out from the same path as the base echo router
	fileRoutes := er.echo.Group("")

	// For serving static resources such as images, javascript, css etc
	// The static directory is respect to the root directory of the application.
	// The separate file systems are for fine grained control so that not any other
	// file stored inside the static directory can be directly accessed.
	tmplJsFS := os.DirFS("static/js")
	tmplCssFS := os.DirFS("static/css")
	tmplImgFS := os.DirFS("static/img")

	// setting up to serve static files
	fileRoutes.StaticFS("/js", tmplJsFS)
	fileRoutes.StaticFS("/css", tmplCssFS)
	fileRoutes.StaticFS("/img", tmplImgFS)

	// Setting up the URLs to serve user facing UIs.
	// Routes in which static HTML is served.
	fileRoutes.GET("/login", controllers.DisplayLoginPage, middlewares.WithAlreadyAuthenticated)
	fileRoutes.GET("/signup", controllers.DisplaySignupPage, middlewares.WithAlreadyAuthenticated)

	// Routes which internally renders the HTML from templates.
	fileRoutes.GET("/helloworld", controllers.RenderHelloWorldPage, middlewares.WithAuthentication)
	fileRoutes.GET("/home", controllers.RenderHomePage, middlewares.WithAuthentication)

}

// Registers routes concerned with the API endpoints.
func (er *EchoRouter) registerApiRoutes() {
	apiV1 := er.echo.Group("/api/v1")

	apiV1.GET("/posts", controllers.ReturnAllPosts)
	apiV1.GET("/posts/:username/:postname", controllers.ReturnSinglePost)

	authApiRoute := apiV1.Group("/auth")
	authApiRoute.POST("/register", controllers.SignUp)
	authApiRoute.POST("/login", controllers.LogIn)
	authApiRoute.POST("/refresh", controllers.Refresh)
	authApiRoute.GET("/logout", controllers.LogOut)

	restrictedApiRoute := apiV1.Group("/")
	restrictedApiRoute.Use(middlewares.WithAuthentication)

	restrictedApiRoute.POST("/posts", controllers.CreatePost)
	restrictedApiRoute.DELETE("/posts/{id}", controllers.DeletePost)
	restrictedApiRoute.PUT("/posts/{id}", controllers.UpdatePost)
}

func (er *EchoRouter) Start() {
	address := fmt.Sprintf(":%d", config.GetPort())
	er.echo.Logger.Fatal(er.echo.Start(address))
}
