package routes

import (
	ctl "api-gin/src/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	v1 := api.Group("v1")

	// Auth
	auth := v1.Group("auth")
	auth.POST("/log-in", ctl.UserLogin)

	// Users
	// user := v1.Group("user")
	// user.GET("/all", ctl.GetAllUser)
	// user.GET("/:id", middleware.JWTMiddleware(), ctl.GetByUserID)

	// Pokemon
	pokemon := v1.Group("pokemon")
	pokemon.GET("/", ctl.GetAllPokemon)
	pokemon.GET("/:id", ctl.GetByPokemonID)
	pokemon.POST("/create", ctl.CreatePokemon)
	pokemon.PUT("/update", ctl.UpdatePokemon)
	pokemon.DELETE("/:id", ctl.DeleteByPokemonID)

	return router

}
