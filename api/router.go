package api

import (
	"fmt"

	"devstream.in/pixelated-pipeline/api/controllers"
	"devstream.in/pixelated-pipeline/api/middlewares"
	"devstream.in/pixelated-pipeline/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router interface {
	RegisterRoutes(allowedOrigins []string)
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

func (er *EchoRouter) RegisterRoutes(allowedOrigins []string) {
	er.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderContentLength,
			echo.HeaderLocation,
			echo.HeaderReferrerPolicy,
			echo.HeaderAuthorization,
			echo.HeaderXRealIP,
			echo.HeaderCookie,
			echo.HeaderAcceptEncoding,
		},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.PATCH,
		},
	}))

	// To declare global router wise static directory.
	er.echo.Static("/", "public")

	// For documentation generation for API, currently not required
	er.echo.GET("/api/swagger/*", echoSwagger.WrapHandler)

	er.registerApiRoutes()
}

// Registers routes concerned with the API endpoints.
func (er *EchoRouter) registerApiRoutes() {
	apiV1 := er.echo.Group("/api/v1")
	{
		apiV1.GET("/hello-text", controllers.HelloTextHandler)
		apiV1.GET("/hello-file", controllers.HelloFileHandler)
		apiV1.GET("/hello-json", controllers.HelloJsonHandler)

		apiV1.GET("/posts", controllers.ReturnAllPosts)
		apiV1.GET("/posts/:username/:postname", controllers.ReturnSinglePost)

		// logout routes do not need to check whether the user is authenticated or not
		// it should simply destroy any authentication token present in the users browser
		// and if session is ongoing, shall also terminate the session.
		apiV1.POST("/logout", controllers.LogOut)

		// A route to forcefully logout from the browser.
		apiV1.GET("/logout", controllers.LogOut)
	}

	authApiRoute := apiV1.Group("/auth", middlewares.WithAlreadyAuthenticated)
	{
		authApiRoute.POST("/register", controllers.SignUp)
		authApiRoute.POST("/login", controllers.LogIn)
		authApiRoute.POST("/refresh", controllers.Refresh)
		authApiRoute.GET("/is-authenticated", controllers.IsAuthenticated)
	}

	restrictedApiRoute := apiV1.Group("/", middlewares.WithAuthentication)
	{
		restrictedApiRoute.GET("/users/:username", controllers.ReturnUserData)
		restrictedApiRoute.POST("/posts", controllers.CreatePost)
		restrictedApiRoute.DELETE("/posts/:id", controllers.DeletePost)
		restrictedApiRoute.PUT("/posts/:id", controllers.UpdatePost)
	}
}

func (er *EchoRouter) Start() {
	address := fmt.Sprintf(":%d", config.GetPort())
	er.echo.Logger.Fatal(er.echo.Start(address))
}
