package api

import (
	"fmt"

	"devstream.in/pixelated-pipeline/api/controllers"
	"devstream.in/pixelated-pipeline/api/middlewares"
	"devstream.in/pixelated-pipeline/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// echoSwagger "github.com/swaggo/echo-swagger"
	// _ "github.com/swaggo/echo-swagger/example/docs"
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
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	}))

	// To declare global router wise static directory.
	// er.echo.Static("/", "static")

	// For documentation generation for API, currently not required
	// er.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	er.registerWebRoutes()
	er.registerApiRoutes()
}

// web routes will serve the index files and required assets in case of a SPA,
// hence to support web routes, it should point to the dist directory of the build
// application and serve the files as a static web server.
func (er *EchoRouter) registerWebRoutes() {
	// used to render go templates, not used when SPA or any javascript
	// based frontend client is used to deploy the website.
	// er.echo.Renderer = NewRenderer("views", true)

	// Group made out from the same path as the base echo router
	// fileRoutes := er.echo.Group("")

	// For serving static resources such as images, javascript, css etc
	// The static directory is respect to the root directory of the application.
	// The separate file systems are for fine grained control so that not any other
	// file stored inside the static directory can be directly accessed.
	// distAssetsFS := os.DirFS("web/dist/assets")
	// distImagesFS := os.DirFS("web/dist/images")

	// setting up to serve static files
	// fileRoutes.StaticFS("/assets", distAssetsFS)
	// fileRoutes.StaticFS("/images", distImagesFS)

	// Setting up the URLs to serve user facing UIs i.e. routes in which static HTML is served.
	// Refer to https://sentry.io/answers/why-don-t-react-router-urls-work-when-refreshing-or-writing-manually/
	// Path Matching Order https://echo.labstack.com/docs/routing#path-matching-order
	// Match Any routes are mapped at the end, so will not be a problem for /api/ routes as they will be
	// matched before, then anything fallback shows the index page. The handlerFunc is hardcoded here
	// as only a single static file is served over all the routes, that is the entry point to the website.
	// fileRoutes.GET("/*", func(c echo.Context) error {
	// 	return c.File("web/dist/index.html")
	// })

	// route to check whether api backend working or not
	er.echo.GET("/api/home", func(c echo.Context) error {
		return c.File("home.txt")
	})
}

// Registers routes concerned with the API endpoints.
func (er *EchoRouter) registerApiRoutes() {
	apiV1 := er.echo.Group("/api/v1")

	apiV1.GET("/posts", controllers.ReturnAllPosts)
	apiV1.GET("/posts/:username/:postname", controllers.ReturnSinglePost)

	// logout routes do not need to check whether the user is authenticated or not
	// it should simply destroy any authentication token present in the users browser
	// and if session is ongoing, shall also terminate the session.
	apiV1.POST("/logout", controllers.LogOut)

	// A route to forcefully logout from the browser.
	apiV1.GET("/logout", controllers.LogOut)

	authApiRoute := apiV1.Group("/auth", middlewares.WithAlreadyAuthenticated)
	authApiRoute.POST("/register", controllers.SignUp)
	authApiRoute.POST("/login", controllers.LogIn)
	authApiRoute.POST("/refresh", controllers.Refresh)
	authApiRoute.GET("/is-authenticated", controllers.IsAuthenticated)

	restrictedApiRoute := apiV1.Group("/")
	restrictedApiRoute.Use(middlewares.WithAuthentication)

	restrictedApiRoute.GET("/users/:username", controllers.ReturnUserData)
	restrictedApiRoute.POST("/posts", controllers.CreatePost)
	restrictedApiRoute.DELETE("/posts/:id", controllers.DeletePost)
	restrictedApiRoute.PUT("/posts/:id", controllers.UpdatePost)
}

func (er *EchoRouter) Start() {
	address := fmt.Sprintf(":%d", config.GetPort())
	er.echo.Logger.Fatal(er.echo.Start(address))
}
