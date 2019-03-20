package main

import (
	"echo-framework/Handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)
func main() {
	app := echo.New()

	app.Logger.SetLevel(log.ERROR)
	app.Use(middleware.Logger())
	app.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(Handlers.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for register and login
			if c.Path() == "/login" || c.Path() == "/register" {
				return true
			}
			return false
		},
	}))

	// Database connection
	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		app.Logger.Fatal(err)
	}

	if err = db.Copy().DB("FlutterEgypt").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	// Initialize Controllers
	h := &Handlers.Handler{DB: db}

	// Routes
	app.POST("/register", h.SignUp)
	app.POST("/login", h.Login)
	app.POST("/posts", h.CreatePost)
	app.GET("/posts", h.FetchPost)
	app.Static("/swagger","swagger-ui/dist")
	app.Logger.Fatal(app.Start(":3000"))
}